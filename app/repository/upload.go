package repository

import (
	"avenue/app/model"
	"context"
	"sync"

	"github.com/google/uuid"
)

type RepositoryResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

type IRepository interface {
	Upload(ctx context.Context, file *model.Upload) (*RepositoryResponse, error)
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
	file.ID = uuid.New().String()

	repository.store.Store(file.ID, file)

	response := &RepositoryResponse{
		Success: true,
		Id:      file.ID,
	}

	return response, nil
}
