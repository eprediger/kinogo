package infrastructure

import (
	logger "application/ports/logging"
	repos "application/ports/repositories"
	"context"
	domain "domain/model"
)

type memorySourcesRepo struct {
	log     logger.Logger
	sources []domain.Source
}

// NewSourcesRepository returns an instance of SourcesRepository
func NewSourcesRepository(log logger.Logger) repos.SourcesRepository {
	return &memorySourcesRepo{
		log:     log,
		sources: []domain.Source{},
	}
}

func (r *memorySourcesRepo) GetAll(ctx context.Context) []domain.Source {
	r.log.Info(ctx, "Sources successfully found")

	return r.sources
}

func (r *memorySourcesRepo) Save(ctx context.Context, newSource *domain.Source) *domain.Source {
	r.sources = append(r.sources, *newSource)
	r.log.Info(ctx, "Source successfully created")

	return newSource
}
