# Use the golang:alpine image as the base image for building the application
FROM golang:alpine as builder

# Set the working directory to /build
WORKDIR /build

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire application source code to the working directory
COPY . .

# Set environment variables for the build process
# CGO_ENABLED=0: Disables CGO (C Go) support, which is useful for building a statically-linked binary without C dependencies
# GOOS=linux: Specifies the target operating system
# GOARCH=amd64: Specifies the target architecture as 64-bit x86
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the application with specific flags and output the binary to /bin/http
# -ldflags="-s -w": Strips the symbol table and debugging information from the binary to reduce its size
RUN go build -ldflags="-s -w" -o /bin/http ./cmd/http/main.go

# Start a new build stage using the scratch base image
FROM scratch

# Copy the CA certificates from the builder stage to the appropriate location in the scratch image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the compiled binary from the builder stage to the root directory of the scratch image
COPY --from=builder ["/build/bin/http", "/http"]

# Set an environment variable to indicate the production environment
ENV GO_ENV=production

# Specify the command to run when the container starts
CMD ["/http"]

