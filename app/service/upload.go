package service

import (
	"avenue/app/model"
	"avenue/app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type IService interface {
	Upload(ctx context.Context, upload *model.Upload) (*repository.RepositoryResponse, error)
}

type UploadService struct {
	repository repository.IRepository
}

func Execute(repository repository.IRepository) IService {
	service := &UploadService{
		repository: repository,
	}

	return service
}

func (service *UploadService) Upload(ctx context.Context, upload *model.Upload) (*repository.RepositoryResponse, error) {
	location := fmt.Sprintf("%s-%s", uuid.New().String(), upload.FileName)
	RepositoryResponse, err := service.repository.Upload(ctx, &model.Upload{
		Path:     location,
		FileName: upload.FileName,
	})

	return RepositoryResponse, err
}
