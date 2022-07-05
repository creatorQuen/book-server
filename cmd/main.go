package main

import (
	"book-server/app"
	"book-server/database"
	"book-server/infrastructure/handlers"
	"book-server/infrastructure/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	sanityCheck()

	db := database.Connect(&gorm.DB{})
	database.Migrate(db)

	bookRepository := repository.NewBookRepository(db)
	bookService := app.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/book", bookHandler.GetBooks)
	router.GET("/book/:id", bookHandler.GetDetailBook)
	router.POST("/book", bookHandler.CreateBookHandler)
	router.PUT("/book/:id", bookHandler.UpdateBookHandler)
	router.DELETE("/book/:id", bookHandler.DeleteBook)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(router.Run(fmt.Sprintf("%s:%s", address, port)))
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviroment variable not defined...")
	}
}
