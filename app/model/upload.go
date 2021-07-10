package model

import "mime/multipart"

type Upload struct {
	ID       string
	FileName string
	Path     string
	ReqId    string
	Headers  multipart.File
}

type UploadResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}
