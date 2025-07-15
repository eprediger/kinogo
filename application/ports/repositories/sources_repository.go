package application

import (
	"context"
	"domain/model"
)

type SourcesRepository interface {
	GetAll(ctx context.Context) []domain.Source
	Save(ctx context.Context, source *domain.Source) *domain.Source
}
