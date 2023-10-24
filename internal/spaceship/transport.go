package spaceship

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ht "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	"github.com/wndisra/galactic-svc/internal/helpers"
)

func RegisterRoutes(router *httprouter.Router, s Service) {
	opts := []ht.ServerOption{
		ht.ServerErrorEncoder(helpers.EncodeError),
	}

	createHandler := ht.NewServer(
		MakeEndpointCreate(s),
		decodeCreateRequest,
		encodeCreateResponse,
		opts...,
	)

	getByIDHandler := ht.NewServer(
		MakeEndpointGetByID(s),
		decodeGetByIDRequest,
		encodeGetByIDResponse,
		opts...,
	)

	updateHandler := ht.NewServer(
		MakeEndpointUpdate(s),
		decodeUpdateRequest,
		encodeUpdateResponse,
		opts...,
	)

	deleteByIDHandler := ht.NewServer(
		MakeEndpointDeleteByID(s),
		decodeDeleteByIDRequest,
		encodeDeleteByIDResponse,
		opts...,
	)

	getAllHandler := ht.NewServer(
		MakeEndpointGetAll(s),
		decodeGetAllRequest,
		encodeGetAllResponse,
		opts...,
	)

	router.Handler(http.MethodPost, "/spaceship", createHandler)
	router.Handler(http.MethodGet, "/spaceship/:id", getByIDHandler)
	router.Handler(http.MethodPatch, "/spaceship", updateHandler)
	router.Handler(http.MethodDelete, "/spaceship/:id", deleteByIDHandler)
	router.Handler(http.MethodGet, "/spaceship", getAllHandler)
}

type createRequest struct {
	Name      string        `json:"name"`
	Class     string        `json:"class"`
	Crew      int64         `json:"crew"`
	Image     string        `json:"image"`
	Value     float64       `json:"value"`
	Status    string        `json:"status"`
	Armaments []armamentReq `json:"armament"`
}

type armamentReq struct {
	Title string `json:"title"`
	Qty   int    `json:"qty"`
}

type armamentResponse struct {
	Title string `json:"title"`
	Qty   int    `json:"qty"`
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeCreateRequest(): %s", err)
	}

	armaments := make([]armamentReqModel, len(req.Armaments))
	for i, armament := range req.Armaments {
		armaments[i] = armamentReqModel(armament)
	}

	return CreateRequestModel{
		Name:      req.Name,
		Class:     req.Class,
		Crew:      req.Crew,
		Image:     req.Image,
		Value:     req.Value,
		Status:    req.Status,
		Armaments: armaments,
	}, nil
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(CreateResponseModel)
	if !ok {
		return fmt.Errorf("encodeCreateResponse(): failed cast response")
	}

	formatted := formatCreateResponse(res)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(formatted)
}

func decodeGetByIDRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	params := httprouter.ParamsFromContext(ctx)

	idPath := params.ByName("id")
	if idPath == ":id" || idPath == "" {
		return nil, helpers.ErrInvalidPathParam
	}

	id, err := strconv.ParseInt(idPath, 10, 64)
	if err != nil {
		return nil, helpers.ErrInvalidPathParam
	}

	return GetByIDRequestModel{
		ID: id,
	}, nil
}

func encodeGetByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(GetByIDResponseModel)
	if !ok {
		return fmt.Errorf("encodeGetByIDResponse() error: failed to cast response")
	}

	formatted := formatGetByIDResponse(res)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(formatted)
}

type updateRequest struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Class     string        `json:"class"`
	Crew      int64         `json:"crew"`
	Image     string        `json:"image"`
	Value     float64       `json:"value"`
	Status    string        `json:"status"`
	Armaments []armamentReq `json:"armament"`
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("decodeUpdateRequest(): %s", err)
	}

	armaments := make([]armamentReqModel, len(req.Armaments))
	for i, armament := range req.Armaments {
		armaments[i] = armamentReqModel(armament)
	}

	return UpdateRequestModel{
		ID:        req.ID,
		Name:      req.Name,
		Class:     req.Class,
		Crew:      req.Crew,
		Image:     req.Image,
		Value:     req.Value,
		Status:    req.Status,
		Armaments: armaments,
	}, nil
}

func encodeUpdateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(UpdateResponseModel)
	if !ok {
		return fmt.Errorf("encodeUpdateResponse(): failed cast response")
	}

	formatted := formatUpdateResponse(res)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(formatted)
}

func decodeDeleteByIDRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	params := httprouter.ParamsFromContext(ctx)

	idPath := params.ByName("id")
	if idPath == ":id" || idPath == "" {
		return nil, helpers.ErrInvalidPathParam
	}

	id, err := strconv.ParseInt(idPath, 10, 64)
	if err != nil {
		return nil, helpers.ErrInvalidPathParam
	}

	return DeleteByIDRequestModel{
		ID: id,
	}, nil
}

func encodeDeleteByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(DeleteByIDResponseModel)
	if !ok {
		return fmt.Errorf("encodeDeleteByIDResponse() error: failed to cast response")
	}

	formatted := formatDeleteByIDResponse(res)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(formatted)
}

type spaceShipResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func decodeGetAllRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	queryValues := r.URL.Query()
	name := queryValues.Get("name")
	class := queryValues.Get("class")
	status := queryValues.Get("status")

	return GetAllRequestModel{
		Name:   name,
		Class:  class,
		Status: status,
	}, nil
}

func encodeGetAllResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res, ok := response.(GetAllResponseModel)
	if !ok {
		return fmt.Errorf("encodeGetAllResponse() error: failed to cast response")
	}

	formatted := formatGetAllResponse(res)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(formatted)
}
