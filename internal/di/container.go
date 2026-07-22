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
    auth.POST("/sign-up", userHandler.Create)
    auth.POST("/sign-in", userHandler.Login)
  }

  { 
    sc := r.Group("/")
    sc.Use(middleware.AuthMiddleware())
    sc.GET("/users", userHandler.GetUsers)
    sc.DELETE("/users/:id", userHandler.DeleteUsers)
    sc.GET("/users/:id", userHandler.GetUserByID)
    sc.PATCH("/users/:id", userHandler.PatchUser)
    sc.PATCH("/users/:id/picture", userHandler.UploadPictureProfile)
  }
}