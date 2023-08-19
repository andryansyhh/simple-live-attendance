// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"simple-live-attendance/src/config"
	"simple-live-attendance/src/models"
	"simple-live-attendance/src/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()
	db := config.GetDatabase()

	// Migrate the database
	db.AutoMigrate(&models.Attendance{})

	routers.SetupRoutes(router, db)

	port := os.Getenv("app_port")
	router.Run(fmt.Sprintf(":%s", port))
}
