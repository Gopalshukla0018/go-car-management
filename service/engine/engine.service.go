package engine

import (
	"context"
	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/Gopalshukla0018/go-car-management/store"
)

type Service struct {
	store store.Engine
}

func New(store store.Engine) *Service {
	return &Service{store: store}
}

func (s *Service) CreateEngine(ctx context.Context, req models.EngineRequest) (models.Engine, error) {
	// Models already have validation logic, but you can call it here if exposed
	return s.store.CreateEngine(ctx, &req)
}

func (s *Service) GetEngine(ctx context.Context, id string) (models.Engine, error) {
	return s.store.EngineByID(ctx, id)
}