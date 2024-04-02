package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/artur-karunas/pop-up-museum/configs"
	"github.com/artur-karunas/pop-up-museum/internal/handlers"
	"github.com/artur-karunas/pop-up-museum/internal/repository"
	"github.com/artur-karunas/pop-up-museum/internal/services"
	"github.com/artur-karunas/pop-up-museum/pkg/emailhandling"
	"github.com/artur-karunas/pop-up-museum/pkg/imagehandling"
	"github.com/artur-karunas/pop-up-museum/system"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialization of the config.
	config, err := configs.InitConfig()
	if err != nil {
		logrus.Fatalf("error occured while initializing the config: %s", err.Error())
	}
	logrus.Info("Config successfully initialized")

	// Initialization of the env.
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Setting the mode of the gin.
	InitGinMode(config.Server.Mode)

	// Initialization of the db.
	db, err := system.InitDatabaseConnection(system.DBConfig{
		Name:     config.MySQL.Name,
		Password: os.Getenv("DB_PASSWORD"),
		Host:     config.MySQL.Host,
		Port:     config.MySQL.Port,
		DBName:   config.MySQL.DBName,
	})
	if err != nil {
		logrus.Fatalf("error occured while connecting the database: %s", err.Error())
	}
	logrus.Info("Database connection successfully initialized")

	// Initialization of the image handler.
	imageHandler := imagehandling.NewImageHandler(config.Application.Uploads)

	// Initialization of the email handler.
	emailHandler := emailhandling.NewEmailService(os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		config.Email.Host,
		config.Email.Port,
	)

	// Initialization of the dependencies.
	repos := repository.NewRepository(db)
	services := services.NewService(repos,
		os.Getenv("HASH_SALT"),
		os.Getenv("TOKEN_SIGNING_KEY"),
		emailHandler,
		config.Application.Appeal.Subject,
		config.Application.Appeal.Message,
		config.Application.Reservation.Subject,
		config.Application.Appeal.Message,
	)
	handlers := handlers.NewHandler(services,
		imageHandler,
		emailHandler,
		config.Application.Passupdate.Subject,
		config.Application.Passupdate.Message,
	)

	// Initialization of the server.
	srv := new(system.Server)
	go func() {
		_ = srv.Run(config.Server.Port, handlers.InitHandler())
	}()
	logrus.Info("Server successfully initialized")

	// Shutdown logic.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("Server shutting down...")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
	logrus.Info("Server is down")
}

func InitGinMode(mode string) {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
