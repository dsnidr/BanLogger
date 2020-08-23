FROM golang:alpine

# Install gcc
RUN apk add build-base

# Set necessary go env variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Dependency setup
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source into container
COPY . .

# Build application
RUN go build -o BanLogger -i ./cmd/banlogger/main.go

# Copy binary to /var/banlogger
WORKDIR /var/banlogger
RUN cp /build/BanLogger .

# Expose necessary port
EXPOSE 443

# Run when starting container
CMD ["/var/banlogger/BanLogger"]