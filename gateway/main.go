package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/ekkinox/bazel-demo/proto/calculator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting gRPC gateway on :8888 ...")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to bridge gRPC backend: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
