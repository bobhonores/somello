package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bobhonores/somello/track/internal/controller/track"
)

type Handler struct {
	ctrl *track.Controller
}

func New(ctrl *track.Controller) *Handler {
	return &Handler{}
}

func (h *Handler) GetTrackDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, track.ErrNotFound) {
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
