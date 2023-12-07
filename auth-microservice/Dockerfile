# Build Stage
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copy only the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o /app/bin .
RUN chmod +x /app/bin

# Final Stage
FROM alpine:3.14

WORKDIR /app

ENV PORT=3005

EXPOSE ${PORT}

# Copy the built executable from the previous stage
COPY --from=build /app/bin /app/

# Use CMD instead of ENTRYPOINT for flexibility
CMD [ "/app/bin" ]
