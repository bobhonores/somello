package grpc

import (
	"context"

	"github.com/bobhonores/somello/gen"
	"github.com/bobhonores/somello/internal/grpcutil"
	"github.com/bobhonores/somello/pkg/discovery"
	"github.com/bobhonores/somello/rating/pkg/model"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{RecordId: string(recordID), RecordType: string(recordType)})
	if err != nil {
		return 0, err
	}

	return resp.RatingValue, nil
}

func (g *Gateway) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return nil
}
