package entity

type BloodType string

// 'A', 'B', 'O', 'AB'
const (
	BloodTypeA  BloodType = "A型"
	BloodTypeB  BloodType = "B型"
	BloodTypeO  BloodType = "O型"
	BloodTypeAB BloodType = "AB型"
)

var BloodTypeName = map[string]BloodType{
	"BLOOD_TYPE_NULL": "ひみつにする",
	"A":               "A型",
	"B":               "B型",
	"O":               "O型",
	"AB":              "AB型",
}
