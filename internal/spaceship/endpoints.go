package spaceship

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"github.com/wndisra/galactic-svc/internal/entity"
)

type Service interface {
	Create(ctx context.Context, req entity.SpaceShip) error
	GetByID(ctx context.Context, id int64) (entity.SpaceShip, error)
	Update(ctx context.Context, id int64, req entity.SpaceShip) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context, req entity.SpaceShip) ([]entity.SpaceShip, error)
}

type CreateRequestModel struct {
	Name      string
	Class     string
	Crew      int64
	Image     string
	Value     float64
	Status    string
	Armaments []armamentReqModel
}

type armamentReqModel struct {
	Title string
	Qty   int
}

func (r CreateRequestModel) ToEntity() entity.SpaceShip {
	armaments := make([]entity.Armament, len(r.Armaments))
	for i, armament := range r.Armaments {
		armaments[i] = entity.Armament{
			Title: armament.Title,
			Qty:   armament.Qty,
		}
	}

	return entity.SpaceShip{
		Name:      r.Name,
		Class:     r.Class,
		Crew:      r.Crew,
		Image:     r.Image,
		Value:     r.Value,
		Status:    r.Status,
		Armaments: armaments,
	}
}

type CreateResponseModel struct {
	Success bool
}

// @BasePath    /
// Create       godoc
// @Description Create new spaceship.
// @Tags        Spaceship
// @Accept      json
// @Produce     json
// @Param       request body createRequest true "Request body (JSON)"
// @Success     201
// @Failure     500
// @Router      /spaceship [post]
func MakeEndpointCreate(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointCreate(): failed cast request")
		}

		err = s.Create(ctx, req.ToEntity())
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointCreate(): %w", err)
		}

		return CreateResponseModel{
			Success: true,
		}, nil
	}
}

type GetByIDRequestModel struct {
	ID int64
}

type GetByIDResponseModel struct {
	SpaceShip entity.SpaceShip
}

// @BasePath    /
// GetByID      godoc
// @Description Fetch existing spaceship by a specific ID.
// @Tags        Spaceship
// @Produce     json
// @Param       id path string true "Spaceship ID (integer)"
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /spaceship/{id} [get]
func MakeEndpointGetByID(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetByIDRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointGetByID(): failed cast request")
		}

		spaceship, err := s.GetByID(ctx, req.ID)
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointGetByID(): %w", err)
		}

		return GetByIDResponseModel{
			SpaceShip: spaceship,
		}, nil
	}
}

type UpdateRequestModel struct {
	ID        int64
	Name      string
	Class     string
	Crew      int64
	Image     string
	Value     float64
	Status    string
	Armaments []armamentReqModel
}

func (r UpdateRequestModel) ToEntity() entity.SpaceShip {
	armaments := make([]entity.Armament, len(r.Armaments))
	for i, armament := range r.Armaments {
		armaments[i] = entity.Armament{
			Title: armament.Title,
			Qty:   armament.Qty,
		}
	}

	return entity.SpaceShip{
		Name:      r.Name,
		Class:     r.Class,
		Crew:      r.Crew,
		Image:     r.Image,
		Value:     r.Value,
		Status:    r.Status,
		Armaments: armaments,
	}
}

type UpdateResponseModel struct {
	Success bool
}

// @BasePath    /
// Update       godoc
// @Description Update existing spaceship by a specific ID.
// @Tags        Spaceship
// @Accept      json
// @Produce     json
// @Param       id path string true "Spaceship ID (integer)"
// @Param       request body updateRequest true "Request body (JSON)"
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /spaceship/{id} [patch]
func MakeEndpointUpdate(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(UpdateRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointUpdate(): failed cast request")
		}

		err = s.Update(ctx, req.ID, req.ToEntity())
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointUpdate(): %w", err)
		}

		return UpdateResponseModel{
			Success: true,
		}, nil
	}
}

type DeleteByIDRequestModel struct {
	ID int64
}

type DeleteByIDResponseModel struct {
	Success bool
}

// @BasePath    /
// Delete       godoc
// @Description Delete existing spaceship by a specific ID.
// @Tags        Spaceship
// @Produce     json
// @Param       id path string true "Spaceship ID (integer)"
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /spaceship/{id} [delete]
func MakeEndpointDeleteByID(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(DeleteByIDRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointDeleteByID(): failed cast request")
		}

		err = s.Delete(ctx, req.ID)
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointDeleteByID(): %w", err)
		}

		return DeleteByIDResponseModel{
			Success: true,
		}, nil
	}
}

type GetAllRequestModel struct {
	Name   string
	Class  string
	Status string
}

func (r GetAllRequestModel) ToEntity() entity.SpaceShip {
	return entity.SpaceShip{
		Name:   r.Name,
		Class:  r.Class,
		Status: r.Status,
	}
}

type GetAllResponseModel struct {
	SpaceShip []entity.SpaceShip
}

// @BasePath    /
// GetAll       godoc
// @Description Get all spaceships.
// @Tags        Spaceship
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /spaceship [get]
func MakeEndpointGetAll(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetAllRequestModel)
		if !ok {
			return nil, errors.New("MakeEndpointGetAll(): failed cast request")
		}

		spaceships, err := s.GetAll(ctx, req.ToEntity())
		if err != nil {
			return nil, fmt.Errorf("MakeEndpointGetAll(): %w", err)
		}

		return GetAllResponseModel{
			SpaceShip: spaceships,
		}, nil
	}
}
