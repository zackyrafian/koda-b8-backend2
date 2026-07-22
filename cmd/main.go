package main

import (
	"context"
	"fmt"
	"koda-b8-backend1/internal/di"
	"koda-b8-backend1/internal/middleware"
	_ "koda-b8-backend1/docs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend V2
// @version 1.0
// @description API documentation
// @host localhost:9999
// @BasePath /

func main () {
  r := gin.Default()
  r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
  err := godotenv.Load()
  if err != nil { 
    fmt.Print("Failed Load .ENV")
  }
  r.Static("/uploads", "./uploads")
  r.Use(middleware.Cors())
  conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
  di.Register(r, conn)
  r.Run(":9999")
}