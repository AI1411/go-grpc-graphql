package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefecture(t *testing.T) {
	t.Run("prefecture string", func(t *testing.T) {
		assert.Equal(t, "ひみつにする", PrefectureSecret.String())
		assert.Equal(t, "北海道", PrefectureHokkaido.String())
		assert.Equal(t, "青森県", PrefectureAomori.String())
		// ... 他の都道府県も同様に追加 ...
		assert.Equal(t, "沖縄県", PrefectureOkinawa.String())
		assert.Equal(t, "海外", PrefectureOversea.String())
	})

	t.Run("prefecture name", func(t *testing.T) {
		assert.Equal(t, "ひみつにする", PrefectureName["PREFECTURE_NULL"].String())
		assert.Equal(t, "北海道", PrefectureName["HOKKAIDO"].String())
		assert.Equal(t, "青森県", PrefectureName["AOMORI"].String())
		// ... 他の都道府県も同様に追加 ...
		assert.Equal(t, "沖縄県", PrefectureName["OKINAWA"].String())
		assert.Equal(t, "海外", PrefectureName["OVERSEAS"].String())
	})

	t.Run("prefecture value", func(t *testing.T) {
		assert.Equal(t, "PREFECTURE_NULL", PrefectureValue["ひみつにする"].String())
		assert.Equal(t, "HOKKAIDO", PrefectureValue["北海道"].String())
		assert.Equal(t, "AOMORI", PrefectureValue["青森県"].String())
		// ... 他の都道府県も同様に追加 ...
		assert.Equal(t, "OKINAWA", PrefectureValue["沖縄県"].String())
		assert.Equal(t, "OVERSEAS", PrefectureValue["海外"].String())
	})
}
