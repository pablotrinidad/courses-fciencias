package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	spb "crawler/proto"
	service "crawler/service"
)

const port uint8 = 8000

func main() {
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
