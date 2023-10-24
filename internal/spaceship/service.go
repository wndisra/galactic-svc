package spaceship

import (
	"context"

	"github.com/go-kit/log"

	"github.com/wndisra/galactic-svc/internal/entity"
	"github.com/wndisra/galactic-svc/internal/helpers"
)

type SpaceShipRepository interface {
	Insert(ctx context.Context, req entity.SpaceShip) error
	GetByID(ctx context.Context, id int64) (entity.SpaceShip, error)
	Update(ctx context.Context, id int64, req entity.SpaceShip) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context, req entity.SpaceShip) ([]entity.SpaceShip, error)
	DeleteArmaments(ctx context.Context, spaceshipID int64) error
}

type service struct {
	repo   SpaceShipRepository
	logger log.Logger
}

func NewService(repo SpaceShipRepository, logger log.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) Create(ctx context.Context, req entity.SpaceShip) error {
	err := s.repo.Insert(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id int64) (entity.SpaceShip, error) {
	spaceship, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return entity.SpaceShip{}, err
	}

	if spaceship.ID == 0 {
		return entity.SpaceShip{}, helpers.ErrBadRequest
	}

	return spaceship, nil
}

func (s *service) Update(ctx context.Context, id int64, req entity.SpaceShip) error {
	spaceship, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if spaceship.ID == 0 {
		return helpers.ErrBadRequest
	}

	err = s.repo.DeleteArmaments(ctx, id)
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, id, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	spaceship, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if spaceship.ID == 0 {
		return helpers.ErrBadRequest
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context, req entity.SpaceShip) ([]entity.SpaceShip, error) {
	return s.repo.GetAll(ctx, req)
}
