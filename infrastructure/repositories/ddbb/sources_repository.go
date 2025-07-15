package infrastructure

import (
	logger "application/ports/logging"
	repos "application/ports/repositories"
	"context"
	domain "domain/model"
)

type SourcesRepository struct {
	memorySourcesRepository repos.SourcesRepository
}

type memorySourcesRepo struct {
	logger  logger.Logger
	sources []domain.Source
}

func NewSourcesRepository(logger logger.Logger) repos.SourcesRepository {
	return &memorySourcesRepo{
		logger:  logger,
		sources: []domain.Source{},
	}
}

func (r *memorySourcesRepo) GetAll(ctx context.Context) []domain.Source {
	r.logger.Info(ctx, "Sources successfully found")
	return r.sources
}

func (r *memorySourcesRepo) Save(ctx context.Context, newSource *domain.Source) *domain.Source {
	r.sources = append(r.sources, *newSource)

	r.logger.Info(ctx, "Source successfully created")
	return newSource
}
