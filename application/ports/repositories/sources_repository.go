package application

import (
	"context"
	"domain/model"
)

// SourcesRepository defines the methods to interact with sources data
type SourcesRepository interface {
	// GetById(ctx context.Context, id string) domain.Source
	GetAll(ctx context.Context) []domain.Source
	Save(ctx context.Context, source *domain.Source) *domain.Source
}
