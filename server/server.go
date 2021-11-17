package main

import (
	"context"
	"database/sql"
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

func (*server) MfRead(ctx context.Context, req *pb.MfRequest) (*pb.MfResponse, error) {
	client, err := sql.Open("mysql", "root@/mf?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	cr := crawlingrepository.NewCrawlingRead(client)
	log.Println("DB read start. user input:", req.UserInput.UserId)

	user, err := cr.UserRead(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(user)

	resultBanks, err := cr.BankRead(ctx, req, user[0].OfficeName, req.StartDay, req.LastDay)
	if err != nil {
		return nil, err
	}
	log.Println("test")
	resultCards, err := cr.CardRead(ctx, req, user[0].OfficeName, req.StartDay, req.LastDay)
	if err != nil {
		return nil, err
	}
	user[0].Banks = resultBanks
	user[0].Cards = resultCards
	log.Println("read DB success")

	return &pb.MfResponse{
		Office: user[0],
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
