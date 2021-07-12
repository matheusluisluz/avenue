package model

import "net/textproto"

type Upload struct {
	ID       string
	FileName string
	Path     string
	ReqId    string
	Headers  textproto.MIMEHeader
}

type UploadResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}
