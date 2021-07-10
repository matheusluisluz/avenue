package service

import (
	"avenue/app/model"
	"avenue/app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	Upload(ctx context.Context, upload *model.Upload) (string, error)
}

type UploadService struct {
	repository repository.Repository
}

func Execute(repository repository.Repository) *UploadService {
	service := &UploadService{
		repository: repository,
	}

	return service
}

func (service *UploadService) Upload(ctx context.Context, upload *model.Upload) (repository.RepositoryResponse, error) {
	location := fmt.Sprintf("%s-%s", uuid.New().String(), upload.FileName)
	RepositoryResponse, err := service.repository.Upload(ctx, &model.Upload{
		Path:     location,
		FileName: upload.FileName,
	})

	return RepositoryResponse, err
}
