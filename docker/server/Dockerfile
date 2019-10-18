#
# FoodUnit 3 API server image
#

# Start build stage
FROM golang:1.12-alpine

# Add git in order to clone the repository.
RUN apk add git --no-cache
RUN git clone https://github.com/dominikbraun/foodunit foodunit
WORKDIR foodunit

# Build the application without the symbol table and debug information.
RUN go build -v -ldflags="-s -w" -o ./.target/foodunit-server ./cmd/server/main.go

# Start final stage
FROM alpine:3.10

# Create the main application directory.
RUN mkdir -p /foodunit
WORKDIR /foodunit

# Copy the compiled binary from the build stage into the application directory
# as well as the application configuration file.
COPY --from=0 /go/foodunit/.target/foodunit-server .
ADD app.toml .

ENV PORT 9292

# Run the FoodUnit server.
CMD ./foodunit-server --addr :${PORT}

# Expose the specified port.
EXPOSE ${PORT}