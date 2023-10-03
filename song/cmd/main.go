package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bobhonores/somello/pkg/discovery"
	"github.com/bobhonores/somello/pkg/discovery/consul"
	song "github.com/bobhonores/somello/song/internal/controller/song"
	metadatagateway "github.com/bobhonores/somello/song/internal/gateway/metadata/http"
	ratinggateway "github.com/bobhonores/somello/song/internal/gateway/rating/http"
	httphandler "github.com/bobhonores/somello/song/internal/handler/http"
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
	h := httphandler.New(ctrl)
	http.Handle("/song", http.HandlerFunc(h.GetSongDetails))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
