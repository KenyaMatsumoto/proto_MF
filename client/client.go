package main

import (
	"context"
	"log"
	"os"

	pb "github.com/gyoza-and-beer/proto-MF/crawlingproto"
	"google.golang.org/grpc"
)

func main() {
	port := "localhost:50051"
	cc, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := pb.NewCrawlingServiceClient(cc)

	crawlingWrite(c)
}

func crawlingWrite(c pb.CrawlingServiceClient) {
	req := &pb.UserRequest{
		UserInput: &pb.UserInput{
			UserId: os.Args[1],
		},
		Pass:     os.Args[2],
		SiteKind: 2,
	}

	res, err := c.UserHandler(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Print(res)
}
