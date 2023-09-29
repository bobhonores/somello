package track

import (
	"context"
	"errors"

	metadatamodel "github.com/bobhonores/somello/metadata/pkg/model"
	ratingmodel "github.com/bobhonores/somello/rating/pkg/model"
	"github.com/bobhonores/somello/song/internal/gateway"
	"github.com/bobhonores/somello/song/pkg/model"
)

var ErrNotFound = errors.New("song metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metametadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metametadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.SongDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.SongDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeSong)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
