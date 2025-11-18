package engine

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gopalshukla0018/go-car-management/models"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) EngineByID(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine
	query := `SELECT id, displacement, no_of_cylinders, car_range FROM engines WHERE id = $1`
	row := e.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange)
	if err != nil {
		return engine, err
	}
	return engine, nil
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	newID := uuid.New()
	query := `INSERT INTO engines (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4) RETURNING id`
	
	var id uuid.UUID
	err := e.db.QueryRowContext(ctx, query, newID, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange).Scan(&id)
	if err != nil {
		return models.Engine{}, err
	}

	return models.Engine{
		EngineID:      id,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	query := `UPDATE engines SET displacement=$1, no_of_cylinders=$2, car_range=$3 WHERE id=$4`
	_, err := e.db.ExecContext(ctx, query, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, id)
	if err != nil {
		return models.Engine{}, err
	}
	
	// Return updated object (simplified)
	uid, _ := uuid.Parse(id)
	return models.Engine{
		EngineID:      uid,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}, nil
}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) error {
	query := `DELETE FROM engines WHERE id=$1`
	res, err := e.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("no engine found to delete")
	}
	return nil
}