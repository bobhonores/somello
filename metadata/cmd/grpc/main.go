package main

import (
	"log"
	"net"

	"github.com/bobhonores/somello/gen"
	"github.com/bobhonores/somello/metadata/internal/controller/metadata"
	grpchandler "github.com/bobhonores/somello/metadata/internal/handler/grpc"
	"github.com/bobhonores/somello/metadata/internal/repository/memory"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting the song metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, h)
	srv.Serve(lis)
}
