package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "upsider.crawling/crawlingproto"
)

func main() {
	log.Println("Run MF crawling server")

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCrawlingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
