package model

import (
	"io"
	"mime/multipart"
)

type Upload struct {
	ID       string
	FileName string
	Path     string
	File     io.Reader
	Header   multipart.FileHeader
}

type Chunk struct {
	UploadID string
	Offset   int64 `form:"offset"`
	Limit    int64 `form:"limit" binding:"required"`
}

type UploadResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

type ReadResponse struct {
	Success bool   `json:"success"`
	File    []byte `json:"file"`
}
