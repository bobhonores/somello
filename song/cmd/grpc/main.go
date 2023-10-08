package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/bobhonores/somello/gen"
	"github.com/bobhonores/somello/pkg/discovery"
	"github.com/bobhonores/somello/pkg/discovery/consul"
	song "github.com/bobhonores/somello/song/internal/controller/song"
	metadatagateway "github.com/bobhonores/somello/song/internal/gateway/metadata/grpc"
	ratinggateway "github.com/bobhonores/somello/song/internal/gateway/rating/grpc"
	grpchandler "github.com/bobhonores/somello/song/internal/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "song"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting the song service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := song.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterSongServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
