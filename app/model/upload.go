package model

import "mime/multipart"

type Upload struct {
	ID       string
	FileName string
	Path     string
	ReqId    string
	Headers  multipart.File
}
