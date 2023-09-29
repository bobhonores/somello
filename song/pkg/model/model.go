package model

import "github.com/bobhonores/somello/metadata/pkg/model"

type SongDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}
