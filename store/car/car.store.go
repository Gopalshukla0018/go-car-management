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

func (s Store) GetByID(ctx context.Context, id string) (models.Car, error) {
	var car models.Car 
	query := `Select c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id,e.displacement,e.no_of_cyclinders,e.car_range FROM car c LEft join engine e on c.engine_id = e.id WHERE c.id =$1`
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt, &car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange)
  if err!=nil{
	if err==sql.ErrNoRows{
		return car, nil
	}
	return car, err
  }
  return car, nil
}

func (s Store) GetCarBrand(ctx context.Context, brand string, isEngine bool) {

}
func (s Store) createCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {

}

func (s Store) updateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {

}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {

}
