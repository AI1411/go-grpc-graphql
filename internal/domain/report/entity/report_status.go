package entity

type ReportStatus string

const (
	ReportStatusPendingInt = iota
	ReportStatusRejectedInt
	ReportStatusAcceptedInt

	ReportStatusPending  ReportStatus = "PENDING"
	ReportStatusRejected ReportStatus = "REJECTED"
	ReportStatusAccepted ReportStatus = "ACCEPTED"
)

var ReportStatusValue = map[ReportStatus]int32{
	ReportStatusPending:  ReportStatusPendingInt,
	ReportStatusRejected: ReportStatusRejectedInt,
	ReportStatusAccepted: ReportStatusAcceptedInt,
}

func (u ReportStatus) Int() int32 {
	return ReportStatusValue[u]
}
