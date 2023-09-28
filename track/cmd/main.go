package main

import (
	"log"
	"net/http"

	"github.com/bobhonores/somello/track/internal/controller/track"
	metadatagateway "github.com/bobhonores/somello/track/internal/gateway/metadata/http"
	ratinggateway "github.com/bobhonores/somello/track/internal/gateway/rating/http"
	httphandler "github.com/bobhonores/somello/track/internal/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := track.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/track", http.HandlerFunc(h.GetTrackDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
