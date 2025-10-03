package models

import "time"

type SummarizeRequest struct {
	Text string `json:"text" binding:"required"`
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
}

type AIRequestLog struct {
	ID           int64     `json:"id"`
	OriginalText string    `json:"original_text"`
	SummaryText  string    `json:"summary_text"`
	Prompt       string    `json:"prompt"`
	CreatedAt    time.Time `json:"created_at"`
}
