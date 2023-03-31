package entity

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	ID             uuid.NullUUID
	ReporterUserID uuid.NullUUID
	ReportedUserID uuid.NullUUID
	ReportedChatID uuid.NullUUID
	Status         ReportStatus
	Reason         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
