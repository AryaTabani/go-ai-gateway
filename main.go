package main

import (
	db "AIRESTAPI/DB"
	"AIRESTAPI/controllers"
	repository "AIRESTAPI/repositories"
	"AIRESTAPI/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment")
	}

	db.InitDB()

	aiRepo := repository.NewAIRequestRepository()
	aiSvc := services.NewAIService(aiRepo)
	aiCtrl := controllers.NewAIController(aiSvc)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		aiRoutes := v1.Group("/ai")
		{
			aiRoutes.POST("/summarize", aiCtrl.HandleSummarize)
		}
	}

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
