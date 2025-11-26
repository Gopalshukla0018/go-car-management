package service

import (
	"context"

	"github.com/Gopalshukla0018/go-car-management/models"
)

type CarService interface {
	GetCar(ctx context.Context, id string) (models.Car, error)
	CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error)
	GetCarsByBrand(ctx context.Context, brand string) ([]models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq models.CarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id string) error
}

type EngineService interface {
	CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error)
	GetEngine(ctx context.Context, id string) (models.Engine, error)
}
