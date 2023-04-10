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

type ReportCount struct {
	ReportedUserID uuid.NullUUID
	ReportCount    int
}

func NewReport(reporterUserID, reportedUserID, reportedChatID uuid.NullUUID, reason string) *Report {
	return &Report{
		ReporterUserID: reporterUserID,
		ReportedUserID: reportedUserID,
		ReportedChatID: reportedChatID,
		Reason:         reason,
	}
}

func NewReportCount(reportedUserID uuid.NullUUID, reportCount int) *ReportCount {
	return &ReportCount{
		ReportedUserID: reportedUserID,
		ReportCount:    reportCount,
	}
}
