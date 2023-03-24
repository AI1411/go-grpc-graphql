package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloodType(t *testing.T) {
	t.Run("blood type test", func(t *testing.T) {
		t.Run("blood type string", func(t *testing.T) {
			assert.Equal(t, "A型", BloodTypeA.String())
			assert.Equal(t, "B型", BloodTypeB.String())
			assert.Equal(t, "O型", BloodTypeO.String())
			assert.Equal(t, "AB型", BloodTypeAB.String())
		})

		t.Run("blood type name", func(t *testing.T) {
			assert.Equal(t, "ひみつにする", BloodTypeName["BLOOD_TYPE_NULL"].String())
			assert.Equal(t, "A型", BloodTypeName["A"].String())
			assert.Equal(t, "B型", BloodTypeName["B"].String())
			assert.Equal(t, "O型", BloodTypeName["O"].String())
			assert.Equal(t, "AB型", BloodTypeName["AB"].String())
		})

		t.Run("blood type value", func(t *testing.T) {
			assert.Equal(t, "BLOOD_TYPE_NULL", BloodTypeValue["ひみつにする"].String())
			assert.Equal(t, "A", BloodTypeValue["A型"].String())
			assert.Equal(t, "B", BloodTypeValue["B型"].String())
			assert.Equal(t, "O", BloodTypeValue["O型"].String())
			assert.Equal(t, "AB", BloodTypeValue["AB型"].String())
		})
	})
}
