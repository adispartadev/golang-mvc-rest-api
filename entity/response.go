package entity

type Response struct {
	Status     Status              `json:"status"`
	Data       interface{}         `json:"data"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

type TokenResponse struct {
	Status       Status `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
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

func SetTokenResponse(code int, message string, token string, refreshToken string) *TokenResponse {
	res := &TokenResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Token:        token,
		RefreshToken: refreshToken,
	}
	return res
}

func SetPaginationResponse(code int, message string, data interface{}, pagination *ResponsePagination) *Response {
	res := &Response{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data:       data,
		Pagination: pagination,
	}
	return res
}
