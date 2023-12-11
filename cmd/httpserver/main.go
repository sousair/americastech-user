package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	gorm_models "github.com/sousair/americastech-user/internal/infra/database/models"
	http_handlers "github.com/sousair/americastech-user/internal/presentation/http/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("/.env")

	if err != nil {
		panic(err)
	}

	postgresConnectionURL := os.Getenv("POSTGRES_CONNECTION_URL")

	db, err := gorm.Open(postgres.Open(postgresConnectionURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&gorm_models.User{})

	e := echo.New()

	e.POST("/users", http_handlers.CreateUserHandler(db))

	e.Logger.Fatal(e.Start(":8080"))
}
