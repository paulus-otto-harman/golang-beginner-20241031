package util

import "20241031/class/1/model"

func BuildResponse(statusCode int, message string, data interface{}) *model.Response {
	return &model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
