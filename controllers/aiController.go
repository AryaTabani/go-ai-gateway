package controllers

import (
	"AIRESTAPI/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	service services.AIService
}

func NewAIController(service services.AIService) *AIController {
	return &AIController{service: service}
}

type SummarizeRequest struct {
	Text   string `json:"text" binding:"required"`
	Prompt string `json:"prompt"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
}

func (ac *AIController) HandleSummarize(c *gin.Context) {
	var req SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	summary, err := ac.service.SummarizeText(c.Request.Context(), req.Text, req.Prompt)
	if err != nil {
		log.Printf("Error from service layer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	c.JSON(http.StatusOK, SummarizeResponse{Summary: summary})
}
