CREATE TABLE IF NOT EXISTS users (
    id  BIGINT PRIMARY KEY,
    email VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR NOT NULL, 
    created_at TIMESTAMPZ DEFAULT now() NOT NULL, 
    updated_at TIMESTAMPZ DEFAULT now() NOT NULL 
);


-- define trigger for update_at field 
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';


CREATE TRIGGER update_timestamp_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();