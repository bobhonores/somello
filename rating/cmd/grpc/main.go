package main

import (
	"log"
	"net"

	"github.com/bobhonores/somello/gen"
	"github.com/bobhonores/somello/rating/internal/controller/rating"
	grpchandler "github.com/bobhonores/somello/rating/internal/handler/grpc"
	"github.com/bobhonores/somello/rating/internal/repository/memory"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterRatingServiceServer(srv, h)
	srv.Serve(lis)
}
