package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

	doCalculate(c, 5, 5, 0)
	doCalculate(c, 6, 5, 1)
	doCalculate(c, 5, 5, 2)
	doCalculate(c, 15, 5, 3)
}
