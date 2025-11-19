package car

import (
	"context"
	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/Gopalshukla0018/go-car-management/store"
)

type Service struct {
	store store.Car
}

func New(store store.Car) *Service {
	return &Service{store: store}
}

func (s *Service) GetCar(ctx context.Context, id string) (models.Car, error) {
	return s.store.GetByID(ctx, id)
}

func (s *Service) CreateCar(ctx context.Context, req models.CarRequest) (models.Car, error) {
	return s.store.CreateCar(ctx, &req)
}

func (s *Service) GetCarsByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	return s.store.GetCarBrand(ctx, brand)
}