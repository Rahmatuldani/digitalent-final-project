package main

import (
	"github.com/Rahmatuldani/digitalent-project/config"
	"github.com/Rahmatuldani/digitalent-project/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	docs "github.com/Rahmatuldani/digitalent-project/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description Server API for MyGram app
// @host localhost:5000
// @BasePath /api/v1

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	config.DBConnect()
	db := config.GetDB()
	validate := validator.New()
	
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	baseRouter := r.Group("/api/v1")
	routers.UserRoutes(db, validate, baseRouter)
	routers.PhotoRoutes(db, validate, baseRouter)
	routers.CommentRoutes(db, validate, baseRouter)
	
	r.Run(":5000")
}