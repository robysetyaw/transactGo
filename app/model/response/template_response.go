package response

type Response struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	StatusData    interface{} `json:"status_data"`
	ErrorMessage  string      `json:"error_message,omitempty"`
}

func NewResponse(code int, message string, data interface{}, error string) Response {
	return Response{
		StatusCode:    code,
		StatusMessage: message,
		StatusData:    data,
		ErrorMessage:  error,
	}
}