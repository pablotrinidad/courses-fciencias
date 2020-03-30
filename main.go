package main

import (
	"log"
	"net"

	"fciencias/crawler"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	crawler.RegisterFCCrawlerServer(s, crawler.NewFCCrawlerServiceImpl())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
