package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/gyoza-and-beer/proto-MF/crawlingproto"
	"github.com/gyoza-and-beer/proto-MF/crawlingrepository"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) UserHandler(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	newC := crawlingrepository.NewCrawling()
	today := time.Now()

	log.Println("crawling start. user input:", req.UserInput.UserId)

	user, banks, details, err := newC.Crawling(req.Pass, req.UserInput)
	if err != nil {
		return &pb.UserResponse{
			IsSuccess: false,
		}, err
	}
	log.Println("crawling success. save DB")

	db := crawlingrepository.NewDatabase()
	if err := db.UserCreate(user, req.UserInput.UserId, &today); err != nil {
		return &pb.UserResponse{
			IsSuccess: false,
		}, err
	}
	if err := db.BankCreate(req.UserInput.UserId, banks, &today); err != nil {
		return &pb.UserResponse{
			IsSuccess: false,
		}, err
	}

	if err := db.DetailCreate(req.UserInput.UserId, details, &today); err != nil {
		return &pb.UserResponse{
			IsSuccess: false,
		}, err
	}

	log.Println("save DB success.")

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
