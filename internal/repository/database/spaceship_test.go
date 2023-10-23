package database

import (
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/wndisra/galactic-svc/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Failed to create sqlmock connection")
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open gorm db session")
	}

	return gormDB, mock
}

func setupMockLogger() kitlog.Logger {
	return kitlog.NewNopLogger()
}

func TestNewRepository(t *testing.T) {
	logger := setupMockLogger()

	mockDB, _ := setupMockDB()
	expected := &repository{
		db:     mockDB,
		logger: logger,
	}

	got := NewRepository(mockDB, logger)
	assert.NotNil(t, got)
	assert.Equal(t, expected, got)
}

func TestRepository_Insert(t *testing.T) {
	query := "INSERT INTO `space_ships` (`created_at`,`updated_at`,`deleted_at`,`name`,`class`,`crew`,`image`,`value`,`status`) VALUES (?,?,?,?,?,?,?,?,?)"
	armamentQuery := "INSERT INTO `armaments` (`created_at`,`updated_at`,`deleted_at`,`title`,`qty`,`space_ship_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `space_ship_id`=VALUES(`space_ship_id`)"

	tests := []struct {
		name    string
		param   entity.SpaceShip
		mocks   func(mock sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "Got error in Gorm query, should return non-nil error",
			param: entity.SpaceShip{
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
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Devastator", "Star Destroyer", 1200, "https://test", 100.99, "Operational").
					WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
			wantErr: assert.AnError,
		},
		{
			name: "Given valid param with no Gorm error, should return nil error",
			param: entity.SpaceShip{
				Name:   "Devastator 2",
				Class:  "Star Destroyer 2",
				Crew:   2200,
				Image:  "https://test",
				Value:  100.99,
				Status: "Operational",
				Armaments: []entity.Armament{
					{
						Title: "Turbo Laser",
						Qty:   10,
					},
				},
			},
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Devastator 2", "Star Destroyer 2", 2200, "https://test", 100.99, "Operational").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec(regexp.QuoteMeta(armamentQuery)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Turbo Laser", 10, sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := setupMockLogger()
			mockDB, mock := setupMockDB()
			sqlDB, _ := mockDB.DB()
			defer sqlDB.Close()

			r := &repository{
				db:     mockDB,
				logger: mockLogger,
			}

			tt.mocks(mock)

			err := r.Insert(context.Background(), tt.param)

			assert.NoError(t, mock.ExpectationsWereMet())
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRepository_GetByID(t *testing.T) {
	query := "SELECT * FROM `space_ships` WHERE id = ? AND `space_ships`.`deleted_at` IS NULL ORDER BY `space_ships`.`id` LIMIT 1"
	armamentQuery := "SELECT * FROM `armaments` WHERE `armaments`.`space_ship_id` = ? AND `armaments`.`deleted_at` IS NULL"

	tests := []struct {
		name    string
		id      int64
		mocks   func(mock sqlmock.Sqlmock)
		want    entity.SpaceShip
		wantErr error
	}{
		{
			name: "Got error in Gorm query, should return empty struct with non-nil error",
			id:   1,
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnError(assert.AnError)
			},
			want:    entity.SpaceShip{},
			wantErr: assert.AnError,
		},
		{
			name: "Given non-existed ID, should return empty struct with nil error",
			id:   2,
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(2).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    entity.SpaceShip{},
			wantErr: nil,
		},
		{
			name: "Given existed ID, should return non-empty struct with nil error",
			id:   3,
			mocks: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "class", "crew", "image", "value", "status"}).
						AddRow(3, "Devastator", "Star Destroyer", 15000, "https://test", 200.99, "Operational"))
				mock.ExpectQuery(regexp.QuoteMeta(armamentQuery)).
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows([]string{"title", "qty", "space_ship_id"}).
						AddRow("Turbo Laser", 10, 3))
			},
			want: entity.SpaceShip{
				Name:   "Devastator",
				Class:  "Star Destroyer",
				Crew:   15000,
				Image:  "https://test",
				Value:  200.99,
				Status: "Operational",
				Armaments: []entity.Armament{
					{
						Title:       "Turbo Laser",
						Qty:         10,
						SpaceShipID: 3,
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := setupMockLogger()
			mockDB, mock := setupMockDB()
			sqlDB, _ := mockDB.DB()
			defer sqlDB.Close()

			r := &repository{
				db:     mockDB,
				logger: mockLogger,
			}

			tt.mocks(mock)

			got, err := r.GetByID(context.Background(), tt.id)

			assert.NoError(t, mock.ExpectationsWereMet())
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Class, got.Class)
			assert.Equal(t, tt.want.Crew, got.Crew)
			assert.Equal(t, tt.want.Image, got.Image)
			assert.Equal(t, tt.want.Value, got.Value)
			assert.Equal(t, tt.want.Status, got.Status)
			assert.Equal(t, tt.want.Armaments, got.Armaments)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRepository_Update(t *testing.T)          {}
func TestRepository_Delete(t *testing.T)          {}
func TestRepository_GetAll(t *testing.T)          {}
func TestRepository_DeleteArmaments(t *testing.T) {}
