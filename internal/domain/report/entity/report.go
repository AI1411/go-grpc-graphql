package entity

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	ID             uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReporterUserID uuid.NullUUID
	ReportedUserID uuid.NullUUID
	ReportedChatID uuid.NullUUID
	Status         ReportStatus
	Reason         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
