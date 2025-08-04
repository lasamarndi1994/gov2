package response

type ResponseDataStruct struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}
type ErrorMessageStruct struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func SuccessMessage(message string) map[string]interface{} {
	response := map[string]interface{}{}
	response["status"] = true
	response["message"] = message

	return response
}

func ResponseData(message string, data interface{}) map[string]interface{} {
	response := map[string]interface{}{}
	response["status"] = true
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}

func ErrorMessage(filed string, message string) interface{} {

	//errors := map[string]interface{}{}
	var errors = map[string]string{
		filed: message,
	}
	var response = ErrorMessageStruct{
		Status:  false,
		Message: "Failed",
		Errors:  errors,
	}
	return response
}

func SuccessResponse(data interface{}) interface{} {
	var response = ResponseDataStruct{
		Status:  true,
		Message: "Success",
		Data:    &data,
	}
	return response
}

func DataBaseError(message string) interface{} {
	// response := map[string]interface{}{}
	// if strings.Contains(message, "Duplicate entry") {
	// 	var filed string

	// 	switch {
	// 	case strings.Contains(message, "uni_users_email"):
	// 		filed = "email"
	// 	case strings.Contains(message, "uni_users_mobile_number"):
	// 	}

	// }
	return message
}
