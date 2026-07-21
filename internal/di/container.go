package di

import (
	"koda-b8-backend1/internal/handler"
	"koda-b8-backend1/internal/repository"
	"koda-b8-backend1/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r *gin.Engine, db *pgxpool.Pool){ 
  userRepo := repository.NewUserRepository(db)
  userService := service.NewUserService(userRepo)
  userHandler := handler.NewUserHandler(userService)

  r.POST("/sign-up", userHandler.Create)
  r.POST("/sign-in", userHandler.Login)
  // r.GET("/users", userHandler.GetUsers)
}