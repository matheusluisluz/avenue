package service

import (
	"context"
	"fmt"

	"avenue/app/model"
	"avenue/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IService interface {
	Upload(ctx context.Context, upload *model.Upload) (*repository.RepositoryResponse, error)
	UploadTest(c *gin.Context, upload *model.Upload) (*repository.RepositoryResponse, error)
	Read(read *model.Chunk) (*model.Upload, error)
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
	id := uuid.New().String()
	location := fmt.Sprintf("%s-%s", id, upload.FileName)
	RepositoryResponse, err := service.repository.Upload(ctx, &model.Upload{
		Path:     location,
		FileName: upload.FileName,
		ID:       id,
	})

	return RepositoryResponse, err
}

func (service *UploadService) UploadTest(c *gin.Context, upload *model.Upload) (*repository.RepositoryResponse, error) {
	id := uuid.New().String()
	dist := fmt.Sprintf("%s-%s", id, upload.FileName)
	// if err := c.SaveUploadedFile(file, dist); err != nil {
	// 	panic(err)
	// }

	RepositoryResponse, err := service.repository.Upload(c.Request.Context(), &model.Upload{
		Path:     dist,
		FileName: upload.FileName,
		ID:       id,
	})

	return RepositoryResponse, err
}

func (service *UploadService) Read(read *model.Chunk) (*model.Upload, error) {
	file, err := service.repository.Read(read)
	return file, err
}
