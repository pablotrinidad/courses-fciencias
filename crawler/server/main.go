package main

import (
	"log"
	"net"

	"crawler/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	service.RegisterFCCrawlerServer(s, service.NewFCCrawlerServiceImpl())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
