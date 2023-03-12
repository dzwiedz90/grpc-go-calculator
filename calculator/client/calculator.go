package main

import (
	"context"
	"log"

	pb "github.com/dzwiedz90/grpc-go-calculator/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient, ctx context.Context, first_number float64, second_number float64, operation int) {
	log.Printf("doGreet was invoked")
	res, err := c.Calculate(ctx, &pb.CalculatorRequest{
		FirstNumber:   first_number,
		SecondNumber:  second_number,
		MathOperation: pb.MathOperation(operation),
	})

	if err != nil {
		log.Fatalf("Could not calculate %v\n", err)
	}

	log.Printf("Result: %f\n", res.Result)
}
