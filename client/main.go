package main

import (
	"context"
	"log"

	pb "github.com/ekkinox/bazel-demo/proto/calculator"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Starting gRPC client on :50051 ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	c := pb.NewCalculatorServiceClient(conn)

	res, err := c.Calculate(context.Background(), &pb.OperationRequest{
		Operation: &pb.Operation{
			Operator: pb.Operator_OPERATOR_PRODUCT,
			Number1:  2,
			Number2:  3,
		},
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}
