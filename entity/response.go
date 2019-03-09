package entity

type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var EmptyValue = make([]string, 0)

func SetResponse(code int, message string, data interface{}) *Response {
	res := &Response{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
	return res
}
