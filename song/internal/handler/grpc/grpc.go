package grpc

import (
	"context"
	"errors"

	"github.com/bobhonores/somello/gen"
	"github.com/bobhonores/somello/metadata/pkg/model"
	"github.com/bobhonores/somello/song/internal/controller/song"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedSongServiceServer
	ctrl *song.Controller
}

func New(ctrl *song.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetSongDetails(ctx context.Context, req *gen.GetSongDetailsRequest) (*gen.GetSongDetailsResponse, error) {
	if req == nil || req.SongId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.SongId)
	if err != nil && errors.Is(err, song.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetSongDetailsResponse{
		SongDetails: &gen.SongDetails{
			Metadata: model.MetadataToProto(&m.Metadata),
			Rating:   *m.Rating,
		},
	}, nil
}
