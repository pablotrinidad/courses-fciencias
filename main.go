package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cs.RegisterFCCrawlerServer(s, cs.NewFCCrawlerServiceImpl())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
