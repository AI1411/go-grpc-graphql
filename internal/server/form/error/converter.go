package error

func ConvertFieldErrorDescription(defaultMessage string, tag string, field string, param string) string {
	switch tag {
	case "required":
		return field + "は必須です"
	case "uuid4":
		return field + "はUUID(v4)でなければなりません"
	}
	return defaultMessage
}
