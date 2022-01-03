package main

import (
	"todo_app/config"
	"todo_app/internal/todo"

	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowCredentials: true,
		AllowMethods: "*",

	}))
	handler.RegisterRoutes(app)

	config.ConnectDB()
	port := os.Getenv("PORT")
    err := app.Listen(":" + port)

    if err != nil {
        log.Fatal("Error app failed to start")
        panic(err)
    }
}