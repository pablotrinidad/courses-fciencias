package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	spb "crawler/proto"
	service "crawler/service"
)

func main() {
	port, err := strconv.Atoi(getenv("SERVICE_PORT", "8000"))
	if err != nil {
		log.Fatalf("got invalid tcp port %s", getenv("SERVICE_PORT", "8000"))
		os.Exit(1)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	spb.RegisterFCCrawlerServer(s, service.NewFCCrawlerService())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		os.Exit(1)
	}

	log.Printf("Starting gRPC server at %d/tcp", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		os.Exit(1)
	}
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
