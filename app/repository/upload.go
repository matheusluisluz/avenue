package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"avenue/app/model"
)

type RepositoryResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	File    []byte `json:"file"`
}

type IRepository interface {
	Upload(ctx context.Context, file *model.Upload) (*RepositoryResponse, error)
	Read(read *model.Chunk) (*model.Upload, error)
}

type UploadRepository struct {
	store sync.Map
}

func Execute() *UploadRepository {
	return &UploadRepository{
		store: sync.Map{},
	}
}

func (repository *UploadRepository) Upload(ctx context.Context, file *model.Upload) (*RepositoryResponse, error) {
	fmt.Println("file.ID: ", file.ID)
	fmt.Println("file: ", file)
	repository.store.Store(file.ID, file)

	response := &RepositoryResponse{
		Success: true,
		Id:      file.ID,
	}

	return response, nil
}

func (repository *UploadRepository) Read(read *model.Chunk) (*model.Upload, error) {
	fmt.Println("read.UploadID: ", read.UploadID)
	file, ok := repository.store.Load(read.UploadID)
	fmt.Println("file: ", file)
	if !ok {
		return nil, errors.New("File Not Found !!!")
	}

	return file.(*model.Upload), nil
}
