package logggers

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"

	pb "github.com/dzwiedz90/grpc-go-calculator/calculator/proto"
)

func LogRequest(ctx context.Context, r *pb.CalculatorRequest) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("Could not get metadata from context")
	}

	log.Println(md)
	fmt.Printf("\tAuthorization: %s, method: %s\n", md.Get("authorization"), md.Get("method"))

	fmt.Printf("\tRequest with input data: FirstNumber: %f, SecondNumber: %f, MathOperation: %d\n", r.FirstNumber, r.SecondNumber, pb.MathOperation(r.MathOperation))
}
