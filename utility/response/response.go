package response

func SuccessMessage(message string, data *string) map[string]interface{} {
	response := map[string]interface{}{}
	response["status"] = true
	response["message"] = message
	if data != nil {
		response["data"] = *data
	}
	return response

}

func ErrorMessage(filed string, message string) map[string]interface{} {
	response := map[string]interface{}{}

	response["status"] = false
	response["errors"] = map[string]string{
		filed: message,
	}

	return response

}
