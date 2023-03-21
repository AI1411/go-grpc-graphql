package util

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStringToNullUUID(t *testing.T) {
	// 正常なUUID文字列のテスト
	t.Run("UUID string test", func(t *testing.T) {
		t.Run("Valid UUID string", func(t *testing.T) {
			validUUIDString := "550e8400-e29b-41d4-a716-446655440000"
			validUUID, err := uuid.Parse(validUUIDString)
			assert.NoError(t, err, "Parsing valid UUID string should not return an error")

			nullUUID := StringToNullUUID(validUUIDString)
			assert.True(t, nullUUID.Valid, "The NullUUID should be valid")
			assert.Equal(t, validUUID, nullUUID.UUID, "The NullUUID should have the correct UUID value")
		})

		t.Run("Empty UUID string", func(t *testing.T) {
			// 空のUUID文字列のテスト
			emptyUUIDString := ""
			nullUUID := StringToNullUUID(emptyUUIDString)
			assert.False(t, nullUUID.Valid, "The NullUUID should be invalid")

			// 無効なUUID文字列のテスト
			invalidUUIDString := "invalid_uuid_string"
			nullUUID = StringToNullUUID(invalidUUIDString)
			assert.False(t, nullUUID.Valid, "The NullUUID should be invalid")
		})
	})
}

func TestNullUUIDToString(t *testing.T) {
	type args struct {
		u uuid.NullUUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid NullUUID",
			args: args{
				u: uuid.NullUUID{
					Valid: true,
					UUID:  uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				},
			},
			want: "550e8400-e29b-41d4-a716-446655440000",
		},
		{
			name: "Invalid NullUUID",
			args: args{
				u: uuid.NullUUID{
					Valid: false,
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NullUUIDToString(tt.args.u), "NullUUIDToString(%v)", tt.args.u)
		})
	}
}
