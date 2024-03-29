package main

import (
	"log"
	"os"

	dynamicsegmentation "github.com/RCNRC/dynamic_segmentation"
	"github.com/RCNRC/dynamic_segmentation/pkg/handler"
	"github.com/RCNRC/dynamic_segmentation/pkg/repository"
	"github.com/RCNRC/dynamic_segmentation/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	/*
		Заметки:
			1) Порт для постгреса сейчас используется 5436, переводящий в докер в 5432, пользователь postgres, пароль qwerty
			2) Нужно отдельно устанавливать migrate
			3) Для timestamp существуют функции
			4) Подключение к postgres происходит через непонятное `pgx` из _ "github.com/jackc/pgx/v5/stdlib"
	*/
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(dynamicsegmentation.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
