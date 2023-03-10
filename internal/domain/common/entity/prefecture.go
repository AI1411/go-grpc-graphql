package entity

type Prefecture string

const (
	PrefectureSecret    Prefecture = "ひみつにする"
	PrefectureHokkaido  Prefecture = "北海道"
	PrefectureAomori    Prefecture = "青森県"
	PrefectureIwate     Prefecture = "岩手県"
	PrefectureMiyagi    Prefecture = "宮城県"
	PrefectureAkita     Prefecture = "秋田県"
	PrefectureYamagata  Prefecture = "山形県"
	PrefectureFukushima Prefecture = "福島県"
	PrefectureIbaraki   Prefecture = "茨城県"
	PrefectureTochigi   Prefecture = "栃木県"
	PrefectureGunma     Prefecture = "群馬県"
	PrefectureSaitama   Prefecture = "埼玉県"
	PrefectureChiba     Prefecture = "千葉県"
	PrefectureTokyo     Prefecture = "東京都"
	PrefectureKanagawa  Prefecture = "神奈川県"
	PrefectureNiigata   Prefecture = "新潟県"
	PrefectureToyama    Prefecture = "富山県"
	PrefectureIshikawa  Prefecture = "石川県"
	PrefectureFukui     Prefecture = "福井県"
	PrefectureYamanashi Prefecture = "山梨県"
	PrefectureNagano    Prefecture = "長野県"
	PrefectureGifu      Prefecture = "岐阜県"
	PrefectureShizuoka  Prefecture = "静岡県"
	PrefectureAichi     Prefecture = "愛知県"
	PrefectureMie       Prefecture = "三重県"
	PrefectureShiga     Prefecture = "滋賀県"
	PrefectureKyoto     Prefecture = "京都府"
	PrefectureOsaka     Prefecture = "大阪府"
	PrefectureHyogo     Prefecture = "兵庫県"
	PrefectureNara      Prefecture = "奈良県"
	PrefectureWakayama  Prefecture = "和歌山県"
	PrefectureTottori   Prefecture = "鳥取県"
	PrefectureShimane   Prefecture = "島根県"
	PrefectureOkayama   Prefecture = "岡山県"
	PrefectureHiroshima Prefecture = "広島県"
	PrefectureYamaguchi Prefecture = "山口県"
	PrefectureTokushima Prefecture = "徳島県"
	PrefectureKagawa    Prefecture = "香川県"
	PrefectureEhime     Prefecture = "愛媛県"
	PrefectureKochi     Prefecture = "高知県"
	PrefectureFukuoka   Prefecture = "福岡県"
	PrefectureSaga      Prefecture = "佐賀県"
	PrefectureNagasaki  Prefecture = "長崎県"
	PrefectureKumamoto  Prefecture = "熊本県"
	PrefectureOita      Prefecture = "大分県"
	PrefectureMiyazaki  Prefecture = "宮崎県"
	PrefectureKagoshima Prefecture = "鹿児島県"
	PrefectureOkinawa   Prefecture = "沖縄県"
	PrefectureOversea   Prefecture = "海外"
)

var (
	PrefectureName = map[string]Prefecture{
		"PREFECTURE_NULL": "ひみつにする",
		"HOKKAIDO":        "北海道",
		"AOMORI":          "青森県",
		"IWATE":           "岩手県",
		"MIYAGI":          "宮城県",
		"AKITA":           "秋田県",
		"YAMAGATA":        "山形県",
		"FUKUSHIMA":       "福島県",
		"IBARAKI":         "茨城県",
		"TOCHIGI":         "栃木県",
		"GUNMA":           "群馬県",
		"SAITAMA":         "埼玉県",
		"CHIBA":           "千葉県",
		"TOKYO":           "東京都",
		"KANAGAWA":        "神奈川県",
		"NIIGATA":         "新潟県",
		"TOYAMA":          "富山県",
		"ISHIKAWA":        "石川県",
		"FUKUI":           "福井県",
		"YAMANASHI":       "山梨県",
		"NAGANO":          "長野県",
		"GIFU":            "岐阜県",
		"SHIZUOKA":        "静岡県",
		"AICHI":           "愛知県",
		"MIE":             "三重県",
		"SHIGA":           "滋賀県",
		"KYOTO":           "京都府",
		"OSAKA":           "大阪府",
		"HYOGO":           "兵庫県",
		"NARA":            "奈良県",
		"WAKAYAMA":        "和歌山県",
		"TOTTORI":         "鳥取県",
		"SHIMANE":         "島根県",
		"OKAYAMA":         "岡山県",
		"HIROSHIMA":       "広島県",
		"YAMAGUCHI":       "山口県",
		"TOKUSHIMA":       "徳島県",
		"KAGAWA":          "香川県",
		"EHIME":           "愛媛県",
		"KOCHI":           "高知県",
		"FUKUOKA":         "福岡県",
		"SAGA":            "佐賀県",
		"NAGASAKI":        "長崎県",
		"KUMAMOTO":        "熊本県",
		"OITA":            "大分県",
		"MIYAZAKI":        "宮崎県",
		"KAGOSHIMA":       "鹿児島県",
		"OKINAWA":         "沖縄県",
		"OVERSEAS":        "海外",
	}

	PrefectureValue = map[string]Prefecture{
		"ひみつにする": "PREFECTURE_NULL",
		"北海道":    "HOKKAIDO",
		"青森県":    "AOMORI",
		"岩手県":    "IWATE",
		"宮城県":    "MIYAGI",
		"秋田県":    "AKITA",
		"山形県":    "YAMAGATA",
		"福島県":    "FUKUSHIMA",
		"茨城県":    "IBARAKI",
		"栃木県":    "TOCHIGI",
		"群馬県":    "GUNMA",
		"埼玉県":    "SAITAMA",
		"千葉県":    "CHIBA",
		"東京都":    "TOKYO",
		"神奈川県":   "KANAGAWA",
		"新潟県":    "NIIGATA",
		"富山県":    "TOYAMA",
		"石川県":    "ISHIKAWA",
		"福井県":    "FUKUI",
		"山梨県":    "YAMANASHI",
		"長野県":    "NAGANO",
		"岐阜県":    "GIFU",
		"静岡県":    "SHIZUOKA",
		"愛知県":    "AICHI",
		"三重県":    "MIE",
		"滋賀県":    "SHIGA",
		"京都府":    "KYOTO",
		"大阪府":    "OSAKA",
		"兵庫県":    "HYOGO",
		"奈良県":    "NARA",
		"和歌山県":   "WAKAYAMA",
		"鳥取県":    "TOTTORI",
		"島根県":    "SHIMANE",
		"岡山県":    "OKAYAMA",
		"広島県":    "HIROSHIMA",
		"山口県":    "YAMAGUCHI",
		"徳島県":    "TOKUSHIMA",
		"香川県":    "KAGAWA",
		"愛媛県":    "EHIME",
		"高知県":    "KOCHI",
		"福岡県":    "FUKUOKA",
		"佐賀県":    "SAGA",
		"長崎県":    "NAGASAKI",
		"熊本県":    "KUMAMOTO",
		"大分県":    "OITA",
		"宮崎県":    "MIYAZAKI",
		"鹿児島県":   "KAGOSHIMA",
		"沖縄県":    "OKINAWA",
		"海外":     "OVERSEAS",
	}
)

func (p Prefecture) String() string {
	return string(p)
}
