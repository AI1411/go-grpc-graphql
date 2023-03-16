package entity

type UserStatus string

var (
	ActiveUser    = []UserStatus{"ACTIVE", "PREMIUM"}
	NotActiveUser = []UserStatus{"RESIGNED", "BANDED"}
)

// 'ACTIVE', 'RESIGNED', 'BANDED', 'PREMIUM'
const (
	UserStatusActive   UserStatus = "通常会員"
	UserStatusResigned UserStatus = "退会済"
	UserStatusBanded   UserStatus = "アカウント停止"
	UserStatusPremium  UserStatus = "プレミアム"
)

var UserStatusName = map[string]UserStatus{
	"ACTIVE":   "通常会員",
	"RESIGNED": "退会済",
	"BANDED":   "アカウント停止",
	"PREMIUM":  "プレミアム",
}

var UserStatusValue = map[string]UserStatus{
	"通常会員":    "ACTIVE",
	"退会済":     "RESIGNED",
	"アカウント停止": "BANDED",
	"プレミアム":   "PREMIUM",
}

func (u UserStatus) String() string {
	return string(u)
}
