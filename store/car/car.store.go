package car

import (
	"context"
	"database/sql"

	"github.com/Gopalshukla0018/go-car-management/models"
)

type Store struct {
	db *sql.DB
}

func new(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s Store) GetyByID(ctx context.Context, id string) models.Car {

}

func (s Store) GetCarBrand(ctx context.Context, brand string, isEngine bool) {

}
func (s Store) createCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {

}

func (s Store) updateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {

}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {

}
