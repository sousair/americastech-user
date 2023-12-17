package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	app_usecases "github.com/sousair/americastech-user/internal/application/usecases"
	bcrypt_cipher "github.com/sousair/americastech-user/internal/infra/cipher"
	gorm_models "github.com/sousair/americastech-user/internal/infra/database/models"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	jwt_provider "github.com/sousair/americastech-user/internal/infra/jwt"
	http_handlers "github.com/sousair/americastech-user/internal/presentation/http/handlers"
	http_middlewares "github.com/sousair/americastech-user/internal/presentation/http/middlewares"
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

	userAuthMiddleware := http_middlewares.UserAuthMiddleware
	userRepo := gorm_repositories.NewUserRepository(db)

	var (
		userPassCostStr = os.Getenv("USER_PASSWORD_COST")
		userTokenSecret = os.Getenv("USER_TOKEN_SECRET")

		port = os.Getenv("PORT")
	)

	userPassCost, err := strconv.Atoi(userPassCostStr)

	if err != nil {
		panic(err)
	}

	validator := validator.New()

	cipherProvider := bcrypt_cipher.NewCipherProvider(userPassCost)
	jwtProvider := jwt_provider.NewJwtProvider(userTokenSecret)

	createUserUC := app_usecases.NewCreateUserUseCase(userRepo, cipherProvider)
	userSignInUC := app_usecases.NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)
	getUsersUC := app_usecases.NewGetUsersUseCase(userRepo)
	getUserUC := app_usecases.NewGetUserUseCase(userRepo)
	updateUserUC := app_usecases.NewUpdateUserUseCase(userRepo)
	deleteUserUC := app_usecases.NewDeleteUserUseCase(userRepo)

	createUserHandler := http_handlers.NewCreateUserHandler(createUserUC, validator).Handle
	userSignInHandler := http_handlers.NewUserSignInHandler(userSignInUC, validator).Handle
	getUsersHandler := http_handlers.NewGetUsersHandler(getUsersUC).Handle
	getUserHandler := http_handlers.NewGetUserHandler(getUserUC, validator).Handle
	updateUserHandler := http_handlers.NewUpdateUserHandler(updateUserUC, validator).Handle
	deleteUserHandler := http_handlers.NewDeleteUserHandler(deleteUserUC, validator).Handle

	e := echo.New()
	e.POST("/users", createUserHandler)
	e.POST("/users/sign-in", userSignInHandler)

	// // ! This should be an admin route in the future
	e.GET("/users", userAuthMiddleware(getUsersHandler))
	e.GET("/users/:id", userAuthMiddleware(getUserHandler))

	e.PUT("/users/:id", userAuthMiddleware(updateUserHandler))
	// // ! This should be an admin route in the future
	e.DELETE("/users/:id", userAuthMiddleware(deleteUserHandler))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
