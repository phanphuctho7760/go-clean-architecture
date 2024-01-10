package entities

import "go/types"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type errorFormat struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func MakeResponseSuccessFormat(message string, data any) (resp Response) {
	resp = Response{
		Message: message,
		Data:    types.Interface{},
	}
	if data != nil {
		resp.Data = data
	}
	return
}

func MakeResponseErrorFormat(message, code, messageError string) (resp Response) {
	resp = Response{
		Message: message,
		Data: errorFormat{
			Code:    code,
			Message: messageError,
		},
	}
	return
}
