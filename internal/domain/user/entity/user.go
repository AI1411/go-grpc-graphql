package entity

import (
	"time"

	"github.com/AI1411/go-grpc-praphql/internal/domain/common/entity"
)

type User struct {
	Id           string
	Username     string
	Email        string
	Password     string
	Status       UserStatus
	Prefecture   entity.Prefecture
	Introduction string
	BloodType    entity.BloodType
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
