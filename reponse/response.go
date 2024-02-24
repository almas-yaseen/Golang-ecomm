package response

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type SuccessResponse struct {
	Statuscode int    `json:"status_code"`
	Message    string `json:"message"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {

	return Response{

		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

}

func SuccessClientResponse(statusCode int, message string) SuccessResponse {
	return SuccessResponse{
		Statuscode: statusCode,
		Message:    message,
	}
}
