package services

type ResponseMessageStruct struct {
	Message string `json:"message"`
}

func ResponseMessage(message string) ResponseMessageStruct {
	var response ResponseMessageStruct
	response.Message = message
	return response
}
