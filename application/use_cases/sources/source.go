package application

import (
	application "application/ports/repositories"
	"context"
	domain "domain/model"
)

// SourceService defines the business logic for sources
type SourceService interface {
	GetAllSources(ctx context.Context) []domain.Source
	CreateSource(ctx context.Context, source *domain.Source) *domain.Source
}

type sourceService struct {
	sourcesRepo application.SourcesRepository
}

// NewSourceService return a new instance of the SourceService
func NewSourceService(sourcesRepo application.SourcesRepository) SourceService {
	return &sourceService{
		sourcesRepo: sourcesRepo,
	}
}

// GetAllSources returns all existing sources.
func (s *sourceService) GetAllSources(ctx context.Context) []domain.Source {
	return s.sourcesRepo.GetAll(ctx)
}

// CreateSource creates a new source
func (s *sourceService) CreateSource(ctx context.Context, source *domain.Source) *domain.Source {
	return s.sourcesRepo.Save(ctx, source)
}
