package main

import (
	"os"
	"strconv"

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

	userPassCostStr := os.Getenv("USER_PASSWORD_COST")
	userPassCost, err := strconv.Atoi(userPassCostStr)

	userTokenSecret := os.Getenv("USER_TOKEN_SECRET")

	if err != nil {
		panic(err)
	}

	cipherProvider := bcrypt_cipher.NewCipherProvider(userPassCost)
	jwtProvider := jwt_provider.NewJwtProvider(userTokenSecret)

	createUserUC := app_usecases.NewCreateUserUseCase(userRepo, cipherProvider)
	userSignInUC := app_usecases.NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)
	getUsersUC := app_usecases.NewGetUsersUseCase(userRepo)
	getUserUC := app_usecases.NewGetUserUseCase(userRepo)
	updateUserUC := app_usecases.NewUpdateUserUseCase(userRepo)
	deleteUserUC := app_usecases.NewDeleteUserUseCase(userRepo)

	createUserHandler := http_handlers.NewCreateUserHandler(createUserUC).Handle
	userSignInHandler := http_handlers.NewUserSignInHandler(userSignInUC).Handle
	getUsersHandler := http_handlers.NewGetUsersHandler(getUsersUC).Handle
	getUserHandler := http_handlers.NewGetUserHandler(getUserUC).Handle
	updateUserHandler := http_handlers.NewUpdateUserHandler(updateUserUC).Handle
	deleteUserHandler := http_handlers.NewDeleteUserHandler(deleteUserUC).Handle

	e := echo.New()
	e.POST("/users", createUserHandler)
	e.POST("/users/sign-in", userSignInHandler)

	// // ! This should be an admin route in the future
	e.GET("/users", userAuthMiddleware(getUsersHandler))
	e.GET("/users/:id", userAuthMiddleware(getUserHandler))

	e.PUT("/users/:id", userAuthMiddleware(updateUserHandler))
	// // ! This should be an admin route in the future
	e.DELETE("/users/:id", userAuthMiddleware(deleteUserHandler))

	e.Logger.Fatal(e.Start(":8080"))
}
