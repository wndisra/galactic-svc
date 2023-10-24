package spaceship

import (
	"context"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/wndisra/galactic-svc/internal/entity"
	"github.com/wndisra/galactic-svc/internal/helpers"
	mock_repo "github.com/wndisra/galactic-svc/internal/repository/database/mocks"
)

func setupMockLogger() kitlog.Logger {
	return kitlog.NewNopLogger()
}

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLogger := setupMockLogger()
	mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)

	expected := &service{
		repo:   mockRepo,
		logger: mockLogger,
	}

	got := NewService(mockRepo, mockLogger)
	assert.NotNil(t, got)
	assert.Equal(t, expected, got)
}

func TestService_Create(t *testing.T) {
	tests := []struct {
		name    string
		req     entity.SpaceShip
		mocks   func(repo *mock_repo.MockSpaceShipRepository)
		wantErr error
	}{
		{
			name: "Got repo error, should return non-nil error",
			req:  entity.SpaceShip{},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().Insert(context.Background(), entity.SpaceShip{}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got repo success, should return nil error",
			req: entity.SpaceShip{
				Name:   "Devastator",
				Class:  "Star Destroyer",
				Crew:   1200,
				Image:  "https://test",
				Value:  100.99,
				Status: "Operational",
				Armaments: []entity.Armament{
					{
						Title: "Turbo Laser",
						Qty:   60,
					},
				},
			},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().Insert(context.Background(), entity.SpaceShip{
					Name:   "Devastator",
					Class:  "Star Destroyer",
					Crew:   1200,
					Image:  "https://test",
					Value:  100.99,
					Status: "Operational",
					Armaments: []entity.Armament{
						{
							Title: "Turbo Laser",
							Qty:   60,
						},
					},
				}).Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)
			mockLogger := setupMockLogger()

			s := &service{
				repo:   mockRepo,
				logger: mockLogger,
			}

			tt.mocks(mockRepo)

			err := s.Create(context.Background(), tt.req)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_GetByID(t *testing.T) {
	spaceship := entity.SpaceShip{
		Name:   "Devastator",
		Class:  "Star Destroyer",
		Crew:   1200,
		Image:  "https://test",
		Value:  100.99,
		Status: "Operational",
		Armaments: []entity.Armament{
			{
				Title: "Turbo Laser",
				Qty:   60,
			},
		},
	}
	spaceship.ID = 2

	tests := []struct {
		name    string
		req     int64
		mocks   func(repo *mock_repo.MockSpaceShipRepository)
		wants   entity.SpaceShip
		wantErr error
	}{
		{
			name: "Got repo error, should return empty struct and non-nil error",
			req:  1,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(1)).Return(entity.SpaceShip{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got repo success but not found, should return non-empty struct and nil error",
			req:  3,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(3)).Return(entity.SpaceShip{}, nil)
			},
			wantErr: helpers.ErrBadRequest,
		},
		{
			name: "Got repo success, should return non-empty struct and nil error",
			req:  2,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
			},
			wants:   spaceship,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)
			mockLogger := setupMockLogger()

			s := &service{
				repo:   mockRepo,
				logger: mockLogger,
			}

			tt.mocks(mockRepo)

			got, err := s.GetByID(context.Background(), tt.req)
			assert.Equal(t, tt.wants, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_Update(t *testing.T) {
	spaceship := entity.SpaceShip{
		Name:   "Devastator",
		Class:  "Star Destroyer",
		Crew:   1200,
		Image:  "https://test",
		Value:  100.99,
		Status: "Operational",
		Armaments: []entity.Armament{
			{
				Title: "Turbo Laser",
				Qty:   60,
			},
		},
	}
	spaceship.ID = 2

	tests := []struct {
		name    string
		id      int64
		req     entity.SpaceShip
		mocks   func(repo *mock_repo.MockSpaceShipRepository)
		wantErr error
	}{
		{
			name: "Got GetByID() repo error, should return non-nil error",
			id:   1,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(1)).Return(entity.SpaceShip{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got GetByID() repo success but not found, should return non-nil error",
			id:   3,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(3)).Return(entity.SpaceShip{}, nil)
			},
			wantErr: helpers.ErrBadRequest,
		},
		{
			name: "Got GetByID() repo success but DeleteArmaments() repo error, should return non-nil error",
			id:   2,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
				repo.EXPECT().DeleteArmaments(context.Background(), int64(2)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got Update() repo error, should return non-nil error",
			id:   2,
			req:  entity.SpaceShip{},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
				repo.EXPECT().DeleteArmaments(context.Background(), int64(2)).Return(nil)
				repo.EXPECT().Update(context.Background(), int64(2), entity.SpaceShip{}).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got repo success, should return nil error",
			id:   2,
			req: entity.SpaceShip{
				Name:   "Devastator",
				Class:  "Star Destroyer",
				Crew:   1200,
				Image:  "https://test",
				Value:  100.99,
				Status: "Operational",
				Armaments: []entity.Armament{
					{
						Title: "Turbo Laser",
						Qty:   60,
					},
				},
			},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
				repo.EXPECT().DeleteArmaments(context.Background(), int64(2)).Return(nil)
				repo.EXPECT().Update(context.Background(), int64(2), entity.SpaceShip{
					Name:   "Devastator",
					Class:  "Star Destroyer",
					Crew:   1200,
					Image:  "https://test",
					Value:  100.99,
					Status: "Operational",
					Armaments: []entity.Armament{
						{
							Title: "Turbo Laser",
							Qty:   60,
						},
					},
				}).Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)
			mockLogger := setupMockLogger()

			s := &service{
				repo:   mockRepo,
				logger: mockLogger,
			}

			tt.mocks(mockRepo)

			err := s.Update(context.Background(), tt.id, tt.req)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_Delete(t *testing.T) {
	spaceship := entity.SpaceShip{
		Name:   "Devastator",
		Class:  "Star Destroyer",
		Crew:   1200,
		Image:  "https://test",
		Value:  100.99,
		Status: "Operational",
		Armaments: []entity.Armament{
			{
				Title: "Turbo Laser",
				Qty:   60,
			},
		},
	}
	spaceship.ID = 2

	tests := []struct {
		name    string
		id      int64
		mocks   func(repo *mock_repo.MockSpaceShipRepository)
		wantErr error
	}{
		{
			name: "Got GetByID() repo error, should return non-nil error",
			id:   1,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(1)).Return(entity.SpaceShip{}, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got GetByID() repo success but not found, should return non-nil error",
			id:   3,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(3)).Return(entity.SpaceShip{}, nil)
			},
			wantErr: helpers.ErrBadRequest,
		},
		{
			name: "Got Delete() repo error, should return non-nil error",
			id:   2,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
				repo.EXPECT().Delete(context.Background(), int64(2)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "Got repo success, should return nil error",
			id:   2,
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetByID(context.Background(), int64(2)).Return(spaceship, nil)
				repo.EXPECT().Delete(context.Background(), int64(2)).Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)
			mockLogger := setupMockLogger()

			s := &service{
				repo:   mockRepo,
				logger: mockLogger,
			}

			tt.mocks(mockRepo)

			err := s.Delete(context.Background(), tt.id)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_GetAll(t *testing.T) {
	// spaceship := entity.SpaceShip{
	// 	Name:   "Devastator",
	// 	Class:  "Star Destroyer",
	// 	Crew:   1200,
	// 	Image:  "https://test",
	// 	Value:  100.99,
	// 	Status: "Operational",
	// 	Armaments: []entity.Armament{
	// 		{
	// 			Title: "Turbo Laser",
	// 			Qty:   60,
	// 		},
	// 	},
	// }
	// spaceship.ID = 2

	tests := []struct {
		name    string
		req     entity.SpaceShip
		mocks   func(repo *mock_repo.MockSpaceShipRepository)
		wants   []entity.SpaceShip
		wantErr error
	}{
		{
			name: "Got GetAll() repo error, should return empty slice and non-nil error",
			req:  entity.SpaceShip{},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetAll(context.Background(), entity.SpaceShip{}).Return([]entity.SpaceShip{}, assert.AnError)
			},
			wants:   []entity.SpaceShip{},
			wantErr: assert.AnError,
		},
		{
			name: "Got repo success, should return non-empty slice and nil error",
			req: entity.SpaceShip{
				Name:   "Devas",
				Class:  "Star Destroyer",
				Status: "Operational",
			},
			mocks: func(repo *mock_repo.MockSpaceShipRepository) {
				repo.EXPECT().GetAll(context.Background(), entity.SpaceShip{
					Name:   "Devas",
					Class:  "Star Destroyer",
					Status: "Operational",
				}).Return([]entity.SpaceShip{}, nil)
			},
			wants:   []entity.SpaceShip{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_repo.NewMockSpaceShipRepository(ctrl)
			mockLogger := setupMockLogger()

			s := &service{
				repo:   mockRepo,
				logger: mockLogger,
			}

			tt.mocks(mockRepo)

			got, err := s.GetAll(context.Background(), tt.req)
			assert.Equal(t, tt.wants, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
