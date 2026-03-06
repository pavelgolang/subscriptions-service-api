package dto

import (
	"github.com/google/uuid"
	"github.com/pavelgolang/subscriptions-service-api/internal/domain"
)

type UpdateSubscriptionDTO struct {
	ID        uint
	Name      *string
	Price     *int
	UserID    *uuid.UUID
	StartDate *domain.Date
	EndDate   *domain.Date
}
