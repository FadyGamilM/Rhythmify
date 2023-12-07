-- Check if user exists before creating
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT FROM pg_user WHERE usename = 'auth_user') THEN
    CREATE USER auth_user WITH PASSWORD 'auth123';
  END IF;
END $$;

-- Create database if not exists
CREATE DATABASE IF NOT EXISTS auth;

-- Grant privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE auth TO auth_user;

-- Connect to the 'auth' database
\c auth;

-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
    id  BIGSERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL, 
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL 
);

-- Define trigger for the 'updated_at' field 
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Create trigger for the 'users' table
CREATE TRIGGER update_timestamp_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();
