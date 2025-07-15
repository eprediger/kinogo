package application

import (
	application "application/ports/repositories"
	"context"
	domain "domain/model"
)

type SourceService interface {
	GetAllSources(ctx context.Context) []domain.Source
	CreateSource(ctx context.Context, source *domain.Source) *domain.Source
}

type sourceService struct {
	sourcesRepo application.SourcesRepository
}

func NewSourcesService(sourcesRepo application.SourcesRepository) SourceService {
	return &sourceService{
		sourcesRepo: sourcesRepo,
	}
}

// GetAllSources returns all the existing sources.
func (s *sourceService) GetAllSources(ctx context.Context) []domain.Source {
	return s.sourcesRepo.GetAll(ctx)
}

// CreateSource saves a new source
func (s *sourceService) CreateSource(ctx context.Context, source *domain.Source) *domain.Source {
	return s.sourcesRepo.Save(ctx, source)
}
