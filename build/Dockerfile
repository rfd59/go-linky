##### Build stage
FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./dist/ ./...



##### Use a minimal base image for the final build
FROM alpine:latest AS runner

# Set destination for COPY
WORKDIR /app

COPY --from=builder /src/dist ./

ENTRYPOINT ["/app/go-linky"]