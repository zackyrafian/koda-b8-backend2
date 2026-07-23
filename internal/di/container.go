package di

import (
	"koda-b8-backend1/internal/handler"
	"koda-b8-backend1/internal/middleware"
	"koda-b8-backend1/internal/repository"
	"koda-b8-backend1/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r *gin.Engine, db *pgxpool.Pool){ 
  userRepo := repository.NewUserRepository(db)
  userService := service.NewUserService(userRepo)
  userHandler := handler.NewUserHandler(userService)

  {
    auth := r.Group("/auth")
    auth.POST("/sign-up", userHandler.Register)
    auth.POST("/sign-in", userHandler.Login)
  }

  { 
    users := r.Group("/")
    users.Use(middleware.AuthMiddleware())
    users.GET("/users", userHandler.GetUsers)
    users.POST("/users", userHandler.Create)
    users.DELETE("/users/:id", userHandler.DeleteUsers)
    users.GET("/users/:id", userHandler.GetUserByID)
    users.PATCH("/users/:id", userHandler.PatchUser)
    users.PATCH("/users/:id/picture", userHandler.UploadPictureProfile)
  }
}