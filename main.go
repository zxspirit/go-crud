package main

import (
	"awesomeProject/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	e := setupRouter()

	err := e.Run("localhost:8080")
	if err != nil {
	}
}

// creat
func setupRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	repo := controllers.New()

	engine.POST("/user", repo.CreateUser)
	engine.GET("users", repo.GetUsers)
	engine.GET("/user/:id", repo.GetUser)
	engine.PUT("/user/:id", repo.UpdateUser)
	engine.DELETE("/user/:id", repo.DeleteUser)

	fmt.Println("4444444444444444444444444444444444444444444444444444444444444")
	return engine
}
