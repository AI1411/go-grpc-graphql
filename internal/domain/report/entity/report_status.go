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

var ReportStatusName = map[int32]ReportStatus{
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
