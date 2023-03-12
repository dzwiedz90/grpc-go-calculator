package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/dzwiedz90/grpc-go-calculator/calculator/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	header := metadata.New(map[string]string{"authorization": "Bearer dupajasi00p13rdzi$t@sioo", "method": "POST"})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	doCalculate(c, ctx, 5, 5, 0)
	doCalculate(c, ctx, 6, 5, 1)
	doCalculate(c, ctx, 5, 5, 2)
	doCalculate(c, ctx, 15, 5, 3)
}
