package model

import "github.com/bobhonores/somello/metadata/pkg/model"

type TrackDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}
