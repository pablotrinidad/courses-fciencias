############### Builder stage
FROM golang:1.14-alpine as builder

# Move source
WORKDIR /go/src/crawler
COPY . .
WORKDIR /go/src/crawler/server

# Fetch dependencies
RUN go mod download
RUN go mod verify

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/service
RUN ls /go/bin

############### Production stage
FROM scratch
COPY --from=builder /go/bin/service /go/bin/service
ENTRYPOINT ["/go/bin/service"]
