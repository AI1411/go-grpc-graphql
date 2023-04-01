package entity

type ReportStatus string

const (
	ReportStatusNullInt = iota
	ReportStatusPendingInt
	ReportStatusRejectedInt
	ReportStatusAcceptedInt

	ReportStatusNull     ReportStatus = "NULL	"
	ReportStatusPending  ReportStatus = "PENDING"
	ReportStatusRejected ReportStatus = "REJECTED"
	ReportStatusAccepted ReportStatus = "ACCEPTED"
)

var ReportStatusValue = map[ReportStatus]int32{
	ReportStatusNull:     ReportStatusNullInt,
	ReportStatusPending:  ReportStatusPendingInt,
	ReportStatusRejected: ReportStatusRejectedInt,
	ReportStatusAccepted: ReportStatusAcceptedInt,
}

var ReportStatusName = map[int32]ReportStatus{
	ReportStatusNullInt:     ReportStatusNull,
	ReportStatusPendingInt:  ReportStatusPending,
	ReportStatusRejectedInt: ReportStatusRejected,
	ReportStatusAcceptedInt: ReportStatusAccepted,
}

func (u ReportStatus) Int() int32 {
	return ReportStatusValue[u]
}

func (u ReportStatus) String() ReportStatus {
	return ReportStatusName[u.Int()]
}
