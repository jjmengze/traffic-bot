package v1alpha1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SearchService struct{}

func (s *SearchService) SearchTrain(ctx context.Context, r *SearchTrainRequest) (*SearchTrainResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
	}

	return &SearchTrainResponse{Message: "value"}, nil
}
