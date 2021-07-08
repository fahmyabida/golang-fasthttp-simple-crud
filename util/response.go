package util

func SuccessWithData(message string, data interface{}, totalCount *int) map[string]interface{} {
	var hMap = make(map[string]interface{})
	hMap["message"] = message
	hMap["data"] = data
	if totalCount != nil {
		hMap["total"] = totalCount
	}
	return hMap
}

func Success(message string) map[string]interface{} {
	var hMap = make(map[string]interface{})
	hMap["message"] = message
	return hMap
}

func Error(errorCode, description string) map[string]interface{} {
	var hMap = make(map[string]interface{})
	hMap["code"] = errorCode
	hMap["description"] = description
	return hMap
}
