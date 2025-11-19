package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s Store) GetByID(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	// Fixed types in your query (LEft -> LEFT, cyclinders -> cylinders)
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.created_at, c.updated_at, 
	                 e.id, e.displacement, e.no_of_cylinders, e.car_range 
	          FROM cars c 
	          LEFT JOIN engines e ON c.engine_id = e.id 
	          WHERE c.id = $1`
	
	row := s.db.QueryRowContext(ctx, query, id)
	// Scanning into Car and nested Engine struct
	err := row.Scan(
		&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Price, &car.CreatedAt, &car.UpdatedAt,
		&car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange,
	)

	if err != nil {
		return car, err
	}
	return car, nil
}

func (s Store) GetCarBrand(ctx context.Context, brand string) ([]models.Car, error) {
	var cars []models.Car
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.created_at, c.updated_at, 
	                 e.id, e.displacement, e.no_of_cylinders, e.car_range 
	          FROM cars c 
	          LEFT JOIN engines e ON c.engine_id = e.id 
	          WHERE c.brand = $1`

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(
			&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Price, &car.CreatedAt, &car.UpdatedAt,
			&car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange,
		); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	newID := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	query := `INSERT INTO cars (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	
	err := s.db.QueryRowContext(ctx, query, 
		newID, carReq.Name, carReq.Year, carReq.Brand, carReq.FuelType, carReq.Engine.EngineID, carReq.Price, createdAt, updatedAt,
	).Scan(&newID)

	if err != nil {
		return models.Car{}, err
	}

	// Construct return object
	return models.Car{
		ID:        newID,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine, // Assuming engine struct is fully populated in request or we just return what we have
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	query := `UPDATE cars SET name=$1, year=$2, brand=$3, fuel_type=$4, price=$5, engine_id=$6, updated_at=$7 WHERE id=$8`
	updatedAt := time.Now()
	
	_, err := s.db.ExecContext(ctx, query, 
		carReq.Name, carReq.Year, carReq.Brand, carReq.FuelType, carReq.Price, carReq.Engine.EngineID, updatedAt, id,
	)
	if err != nil {
		return models.Car{}, err
	}

	uid, _ := uuid.Parse(id)
	return models.Car{
		ID:        uid,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine,
		Price:     carReq.Price,
		UpdatedAt: updatedAt,
	}, nil
}

func (s Store) DeleteCar(ctx context.Context, id string) error {
	query := `DELETE FROM cars WHERE id=$1`
	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("no car found")
	}
	return nil
}