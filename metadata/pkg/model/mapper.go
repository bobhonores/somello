package model

import "github.com/bobhonores/somello/gen"

func MetadataToProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id:     m.ID,
		Title:  m.Title,
		Artist: m.Artist,
		Author: m.Author,
	}
}

func MetadataFromProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		ID:     m.Id,
		Title:  m.Title,
		Artist: m.Artist,
		Author: m.Author,
	}
}
