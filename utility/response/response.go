package response

func SuccessMessage(status string, message string) map[string]string {
	response := map[string]string{}
	response["message"] = message
	response["status"] = status

	return response

}

func ErrorMessage(status bool, message string) map[string]interface{} {
	response := map[string]interface{}{}

	response["status"] = false
	response["message"] = "message"

	return response

}
