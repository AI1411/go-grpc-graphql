package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
)

type User struct {
	ID           uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string
	Email        string
	Password     string
	Status       UserStatus
	Prefecture   entity.Prefecture
	Introduction string
	BloodType    entity.BloodType
	ImagePath    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserPassword struct {
	ID                   uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Password             string
	PasswordConfirmation string
}

func NewUser(
	username, email, password string,
	status UserStatus,
	prefecture entity.Prefecture,
	introduction string,
	bloodType entity.BloodType,
	imagePath string,
) *User {
	return &User{
		Username:     username,
		Email:        email,
		Password:     password,
		Status:       status,
		Prefecture:   prefecture,
		Introduction: introduction,
		BloodType:    bloodType,
		ImagePath:    imagePath,
	}
}
