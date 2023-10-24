package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/wndisra/galactic-svc/docs"
	"github.com/wndisra/galactic-svc/internal/entity"
	"github.com/wndisra/galactic-svc/internal/repository/database"
	"github.com/wndisra/galactic-svc/internal/spaceship"
)

// @title Galactic Service APIs
// @description The server APIs documentation for Galactic.
func main() {
	// Init logger
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// Load env
	err := godotenv.Load()
	if err != nil {
		level.Error(logger).Log("msg", "failed to load env")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Init DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		level.Error(logger).Log("msg", "failed to connect to database")
		os.Exit(1)
	}

	// Migrate database
	db.AutoMigrate(&entity.SpaceShip{}, &entity.Armament{})

	dbRepo := database.NewRepository(db, logger)
	spaceShipSvc := spaceship.NewService(dbRepo, logger)

	// Init router
	router := httprouter.New()
	docs.SwaggerInfo.BasePath = "/"

	// Register routes
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "You're alive!")
	})
	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "Pong!")
	})

	// Spaceships routes
	spaceship.RegisterRoutes(router, spaceShipSvc)

	// Swagger documentation
	// TODO: enable for development env, disable for production env
	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpSwagger.WrapHandler(w, r)
	})

	// Listen & serve request
	level.Info(logger).Log("msg", "server started successfully")
	http.ListenAndServe(":3000", router)
}
