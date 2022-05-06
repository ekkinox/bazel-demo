package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ekkinox/bazel-demo/proto/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.CalculatorServiceServer
}

func main() {

	fmt.Println("Starting gRPC server on :50051 ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Calculate(ctx context.Context, in *pb.OperationRequest) (*pb.OperationResponse, error) {

	log.Printf("Calculate called with %v", in)

	return &pb.OperationResponse{
		Result: calculateOperation(in),
	}, nil
}

func calculateOperation(c *pb.OperationRequest) float32 {
	switch c.Operation.Operator {
	case pb.Operator_OPERATOR_SUM:
		return c.Operation.Number1 + c.Operation.Number2
	case pb.Operator_OPERATOR_DIFF:
		return c.Operation.Number1 - c.Operation.Number2
	case pb.Operator_OPERATOR_PRODUCT:
		return c.Operation.Number1 * c.Operation.Number2
	case pb.Operator_OPERATOR_QUOTIENT:
		return c.Operation.Number1 / c.Operation.Number2
	default:
		return 0
	}
}
