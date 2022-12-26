package models

type Response struct {
	Data      interface{} `json:"data"`
	Code      int         `json:"code"`
	Error     bool        `json:"error"`
	ErrorText string      `json:"error_text"`
}
