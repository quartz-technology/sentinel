###
# Stage: builder
# Install dependencies and build sentinel binary.
# Output:
#  /app/sentinel
###
FROM golang:1.23.4-alpine3.21 AS builder

# Define working directory.
WORKDIR /app

# Install build dependencies.
RUN apk add build-base

# Install Golang dependencies.
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy source code.
COPY . .

# Build binary.
RUN go build -o /app/sentinel ./main.go

###
# Stage: app
# Expose sentinel binary entrypoint in a light image.
###
FROM alpine:3.21 AS app

# Define workding directory
WORKDIR /app

# Copy binary from builder stage to alpine image.
COPY --from=builder /app/sentinel /app/sentinel

# Set entrypoint to  binary.
ENTRYPOINT ["/app/sentinel"]
