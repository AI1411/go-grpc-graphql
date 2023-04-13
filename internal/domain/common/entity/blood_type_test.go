package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-grpc-graphql/internal/domain/common/entity"
)

func TestBloodType(t *testing.T) {
	t.Run("blood type test", func(t *testing.T) {
		t.Run("blood type string", func(t *testing.T) {
			assert.Equal(t, "A型", entity.BloodTypeA.String())
			assert.Equal(t, "B型", entity.BloodTypeB.String())
			assert.Equal(t, "O型", entity.BloodTypeO.String())
			assert.Equal(t, "AB型", entity.BloodTypeAB.String())
		})

		t.Run("blood type name", func(t *testing.T) {
			assert.Equal(t, "ひみつにする", entity.BloodTypeName["BLOOD_TYPE_NULL"].String())
			assert.Equal(t, "A型", entity.BloodTypeName["A"].String())
			assert.Equal(t, "B型", entity.BloodTypeName["B"].String())
			assert.Equal(t, "O型", entity.BloodTypeName["O"].String())
			assert.Equal(t, "AB型", entity.BloodTypeName["AB"].String())
		})

		t.Run("blood type value", func(t *testing.T) {
			assert.Equal(t, "BLOOD_TYPE_NULL", entity.BloodTypeValue["ひみつにする"].String())
			assert.Equal(t, "A", entity.BloodTypeValue["A型"].String())
			assert.Equal(t, "B", entity.BloodTypeValue["B型"].String())
			assert.Equal(t, "O", entity.BloodTypeValue["O型"].String())
			assert.Equal(t, "AB", entity.BloodTypeValue["AB型"].String())
		})
	})
}
