package model

// Metadata definess the song metadata
type Metadata struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Author string `json:"author"`
}
