package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/AI1411/go-grpc-praphql/internal/domain/common/entity"
)

type User struct {
	ID           uuid.NullUUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string
	Email        string
	Password     []byte
	Status       UserStatus
	Prefecture   entity.Prefecture
	Introduction string
	BloodType    entity.BloodType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
