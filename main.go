package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/middleware"
	"github.com/tyhafen.ecommerce-go/controllers"
	"github.com/tyhafen.ecommerce-go/database"
	"github.com/tyhafen.ecommerce-go/middleware"
	"github.com/tyhafen.ecommerce-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.User(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

}
