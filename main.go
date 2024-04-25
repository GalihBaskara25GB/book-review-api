package main

import (
	"bookreview/controllers"
	"bookreview/database"
	"bookreview/middleware"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("./.env")
	if err != nil {
		fmt.Println("failed to load env file")
	} else {
		fmt.Println("env file loaded...")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("database connected...")

	database.DbMigrate(DB)

	defer DB.Close()

	listenOn := "localhost:8080"
	hostingProvider := os.Getenv("HOSTING_PROVIDER")
	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", middleware.Auth([]string{"superuser", "author"}), controllers.InsertBook)
	router.PUT("/books/:id", middleware.Auth([]string{"superuser", "author"}), controllers.UpdateBook)
	router.DELETE("/books/:id", middleware.Auth([]string{"superuser", "author"}), controllers.DeleteBook)

	router.GET("/categories", controllers.GetCategories)
	router.GET("/categories/:id", controllers.GetCategory)
	router.POST("/categories", middleware.Auth([]string{"superuser"}), controllers.InsertCategory)
	router.PUT("/categories/:id", middleware.Auth([]string{"superuser"}), controllers.UpdateCategory)
	router.DELETE("/categories/:id", middleware.Auth([]string{"superuser"}), controllers.DeleteCategory)

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.InsertUser)
	router.PUT("/users/:id", middleware.Auth([]string{}), controllers.UpdateUser)
	router.DELETE("/users/:id", middleware.Auth([]string{}), controllers.DeleteUser)

	router.GET("/reviews", controllers.GetReviews)
	router.GET("/reviews/:id", controllers.GetReview)
	router.POST("/reviews", middleware.Auth([]string{"superuser", "reviewer"}), controllers.InsertReview)
	router.PUT("/reviews/:id", middleware.Auth([]string{"superuser", "reviewer"}), controllers.UpdateReview)
	router.DELETE("/reviews/:id", middleware.Auth([]string{"superuser", "reviewer"}), controllers.DeleteReview)

	if hostingProvider == "railway" {
		appHost = "0.0.0.0"
		appPort = os.Getenv("PORT")
	}

	if appHost != "" {
		if appPort != "" {
			listenOn = appHost + ":" + appPort
		} else {
			listenOn = appHost
		}
	}

	router.Run(listenOn)
	fmt.Println("listening on " + listenOn + " ...")
}
