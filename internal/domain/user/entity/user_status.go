package entity

type UserStatus string

// 'ACTIVE', 'RESIGNED', 'BANDED', 'PREMIUM'
const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusResigned UserStatus = "RESIGNED"
	UserStatusBanded   UserStatus = "BANDED"
	UserStatusPremium  UserStatus = "PREMIUM"
)
