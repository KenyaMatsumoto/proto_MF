package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gyoza-and-beer/proto-MF/crawlingproto"
	"github.com/gyoza-and-beer/proto-MF/crawlingrepository"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) UserHandler(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	newC := crawlingrepository.NewCrawling()
	// today := time.Now()

	user, banks, details, err := newC.Crawling(req.Pass, req.UserInput)
	if err != nil {
		return &pb.UserResponse{
			IsSuccess: false,
		}, err
	}

	log.Println(user)
	log.Println(banks)
	log.Println(banks[0])
	log.Println(banks[1])
	log.Println(banks[2])
	log.Println(details)
	return &pb.UserResponse{
		IsSuccess: true,
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCrawlingServiceServer(s, &server{})

	log.Printf("gRPC server listengin on " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
