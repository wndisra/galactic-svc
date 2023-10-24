package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"gorm.io/gorm"

	"github.com/wndisra/galactic-svc/internal/entity"
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Insert(ctx context.Context, req entity.SpaceShip) error {
	result := r.db.Create(&entity.SpaceShip{
		Name:      req.Name,
		Class:     req.Class,
		Crew:      req.Crew,
		Image:     req.Image,
		Value:     req.Value,
		Status:    req.Status,
		Armaments: req.Armaments,
	})

	if result.Error != nil {
		level.Error(r.logger).Log("msg", "database.Insert(): failed to insert to database")
		return result.Error
	}

	return nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (entity.SpaceShip, error) {
	var spaceship entity.SpaceShip

	result := r.db.Preload("Armaments").First(&spaceship, "id = ?", id)

	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		level.Error(r.logger).Log("msg", "database.GetByID(): failed to fetch from database")
		return entity.SpaceShip{}, err
	}

	return spaceship, nil
}

func (r *repository) Update(ctx context.Context, id int64, req entity.SpaceShip) error {
	var entity entity.SpaceShip

	r.db.First(&entity, "id = ?", id)
	entity.Armaments = req.Armaments // ensure armaments data also updated

	result := r.db.Model(&entity).Updates(req)
	err := result.Error
	if err != nil {
		level.Error(r.logger).Log("msg", "database.Update(): failed to update data in database")
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	var model entity.SpaceShip

	result := r.db.Delete(&model, "id = ?", id)
	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		level.Error(r.logger).Log("msg", "database.Delete(): failed to update data in database")
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context, req entity.SpaceShip) ([]entity.SpaceShip, error) {
	var spaceships []entity.SpaceShip

	whereQuery := "name LIKE " + "'%" + req.Name + "%'"

	if req.Class != "" {
		whereQuery += fmt.Sprintf(" AND class = '%s'", req.Class)
	}

	if req.Status != "" {
		whereQuery += fmt.Sprintf(" AND status = '%s'", req.Status)
	}

	result := r.db.Where(whereQuery).Find(&spaceships)
	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.SpaceShip{}, err
	}

	return spaceships, nil
}

func (r *repository) DeleteArmaments(ctx context.Context, spaceshipID int64) error {
	var model entity.Armament

	result := r.db.Delete(&model, "space_ship_id = ?", spaceshipID)
	err := result.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		level.Error(r.logger).Log("msg", "database.DeleteArmaments(): failed to delete armaments data in database")
		return err
	}

	return nil
}
