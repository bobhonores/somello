package main

import (
	"log"
	"net/http"

	track "github.com/bobhonores/somello/song/internal/controller/song"
	metadatagateway "github.com/bobhonores/somello/song/internal/gateway/metadata/http"
	ratinggateway "github.com/bobhonores/somello/song/internal/gateway/rating/http"
	httphandler "github.com/bobhonores/somello/song/internal/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := track.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/track", http.HandlerFunc(h.GetSongDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
