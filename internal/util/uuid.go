package util

import (
	"github.com/google/uuid"
)

// StringToNullUUID convert string to NullUUID
func StringToNullUUID(s string) (u uuid.NullUUID) {
	if s == "" {
		return
	}

	uuID, err := uuid.Parse(s)
	if err != nil {
		return
	}

	u = uuid.NullUUID{
		Valid: true,
		UUID:  uuID,
	}

	return
}

// NullUUIDToString convert NullUUID to string
func NullUUIDToString(u uuid.NullUUID) string {
	if u.Valid {
		return u.UUID.String()
	}

	return ""
}
