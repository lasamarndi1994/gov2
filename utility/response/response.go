package response

func SuccessMessage(message string) map[string]interface{} {
	response := map[string]interface{}{}
	response["status"] = true
	response["message"] = message
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
