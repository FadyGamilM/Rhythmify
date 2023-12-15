# Build Stage
FROM golang:1.21-alpine AS build

WORKDIR /app/auth

# Copy only the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all Go files from the current directory to /app/auth/
COPY . ./

# Build the Go application
RUN go build -o ./bin ./cmd/
RUN chmod +x ./bin

# Final Stage
FROM alpine:3.14

WORKDIR /app/auth

ENV SERVER_PORT=3005
ENV POSTGRES_HOST=0.0.0.0:5432
ENV POSTGRES_USER=auth_user
ENV POSTGRES_PASSWORD=auth123
ENV POSTGRES_SSLMODE=disable
ENV POSTGRES_DB=auth
ENV JWT_SECRET_KEY=just_test_auth_service_in_microservice_app

EXPOSE ${SERVER_PORT}

# Copy the built executable from the previous stage
COPY --from=build /app/auth/bin ./bin

# Use CMD instead of ENTRYPOINT for flexibility
CMD [ "./bin" ]
