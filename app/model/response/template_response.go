package response

type Response struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Message       interface{} `json:"message"`
	StatusData    interface{} `json:"status_data"`
	ErrorMessage  string      `json:"error_message,omitempty"`
}

func NewResponse(code int, status_message string,message interface{}, data interface{}, error string) Response {
	return Response{
		StatusCode:    code,
		StatusMessage: status_message,
		Message:       message,
		StatusData:    data,
		ErrorMessage:  error,
	}
}