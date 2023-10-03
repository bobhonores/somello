package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	song "github.com/bobhonores/somello/song/internal/controller/song"
)

type Handler struct {
	ctrl *song.Controller
}

func New(ctrl *song.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetSongDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, song.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
