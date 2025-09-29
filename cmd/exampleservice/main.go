package main

import (
	"context"
	"log"
	"net"
	"time"

	exampleapiv1 "github.com/usuginus/go-protobuf-builder/gen/go/example/api/v1"
	examplecommonv1 "github.com/usuginus/go-protobuf-builder/gen/go/example/common/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	listenAddr = ":50051"
)

type exampleService struct {
	exampleapiv1.UnimplementedExampleServiceServer
}

func (s *exampleService) GetExample(ctx context.Context, req *exampleapiv1.GetExampleRequest) (*exampleapiv1.GetExampleResponse, error) {
	now := timestamppb.New(time.Now().UTC())

	example := &examplecommonv1.Example{
		Name:        req.GetName(),
		DisplayName: "Sample Example",
		Description: "Minimal in-memory example returned by the demo service.",
		CreateTime:  now,
		UpdateTime:  now,
	}

	return &exampleapiv1.GetExampleResponse{Example: example}, nil
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", listenAddr, err)
	}

	server := grpc.NewServer()
	exampleapiv1.RegisterExampleServiceServer(server, &exampleService{})
	reflection.Register(server)

	log.Printf("ExampleService listening on %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
