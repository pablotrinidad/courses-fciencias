package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc/grpclog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	spb "crawler/proto"
	crawlersvc "crawler/service"
)

const port uint32 = 8000

func init() {
	logger := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(logger)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			durationInterceptor,
		),
	)
	reflection.Register(s)
	spb.RegisterFCCrawlerServer(s, crawlersvc.NewFCCrawlerService())

	grpclog.Infof("Starting gRPC server at %d/tcp", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		os.Exit(1)
	}
}

// durationInterceptor logs the method being called and the duration of the handler execution.
func durationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	h, handlerErr := handler(ctx, req)
	execDuration := time.Since(start)
	if service, method, err := parseMethod(info.FullMethod); err == nil {
		status := "OK"
		if handlerErr != nil {
			status = "ERROR"
		}
		grpclog.Infof("%s.%s (%s): %s", service, method, execDuration, status)
	} else {
		grpclog.Errorf("failed parsing method name %s; %v", info.FullMethod, err)
	}
	return h, handlerErr
}
