package main

import (
	"context"
	"log"

	pb "github.com/dzwiedz90/grpc-go-calculator/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)

	result := 0.0

	switch in.MathOperation {
	case pb.MathOperation_MATH_OPERATION_ADDITION:
		result = in.FirstNumber + in.SecondNumber
	case pb.MathOperation_MATH_OPERATION_SUBSTRACTION:
		result = in.FirstNumber - in.SecondNumber
	case pb.MathOperation_MATH_OPERATION_MULTIPLICATION:
		result = in.FirstNumber * in.SecondNumber
	case pb.MathOperation_MATH_OPERATION_DIVISION:
		result = in.FirstNumber / in.SecondNumber
	}

	return &pb.CalculatorResponse{
		Result: result,
	}, nil
}
