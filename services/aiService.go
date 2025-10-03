package services

import (
	"AIRESTAPI/models"
	repository "AIRESTAPI/repositories"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type AIService interface {
	SummarizeText(ctx context.Context, text string, prompt string) (string, error)
}

type aiServiceImpl struct {
	repo       repository.AIRequestRepository
	httpClient *http.Client
}

func NewAIService(repo repository.AIRequestRepository) AIService {
	return &aiServiceImpl{
		repo:       repo,
		httpClient: &http.Client{},
	}
}

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}
type Content struct {
	Parts []Part `json:"parts"`
}
type Part struct {
	Text string `json:"text"`
}
type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}
type Candidate struct {
	Content Content `json:"content"`
}

func (s *aiServiceImpl) SummarizeText(ctx context.Context, text string, userPrompt string) (string, error) {
	log.Printf("Received text to process: %.50s...", text)

	summary, err := s.callAImodel(ctx, text, userPrompt)
	if err != nil {
		log.Printf("Error calling AI model: %v", err)
		return "", err
	}

	requestLog := &models.AIRequestLog{
		OriginalText: text,
		SummaryText:  summary,
		Prompt:       userPrompt,
	}

	if err := s.repo.CreateLog(ctx, requestLog); err != nil {
		log.Printf("Warning: Failed to log AI request to database: %v", err)
	}

	return summary, nil
}

func (s *aiServiceImpl) callAImodel(ctx context.Context, text string, userPrompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	apiURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=" + apiKey

	finalPrompt := userPrompt
	if finalPrompt == "" {
		finalPrompt = "Summarize the following text in one sentence: "
	}

	fullPrompt := finalPrompt + text

	reqPayload := GeminiRequest{
		Contents: []Content{{Parts: []Part{{Text: fullPrompt}}}},
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to Gemini API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Gemini API returned non-200 status: %d %s", resp.StatusCode, string(bodyBytes))
	}

	var geminiResp GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return "", fmt.Errorf("failed to decode Gemini API response: %w", err)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no content found in Gemini API response")
}
