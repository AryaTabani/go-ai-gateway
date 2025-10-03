package repository

import (
	db "AIRESTAPI/DB"
	"AIRESTAPI/models"
	"context"
)

type AIRequestRepository interface {
	CreateLog(ctx context.Context, logEntry *models.AIRequestLog) error
}

type aiRequestRepositoryImpl struct{}

func NewAIRequestRepository() AIRequestRepository {
	return &aiRequestRepositoryImpl{}
}

func (r *aiRequestRepositoryImpl) CreateLog(ctx context.Context, logEntry *models.AIRequestLog) error {
	query := `INSERT INTO ai_requests (original_text, summary_text, prompt) VALUES (?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, logEntry.OriginalText, logEntry.SummaryText, logEntry.Prompt)
	return err
}
