package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"avenue/app/model"
	"avenue/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IService interface {
	Upload(ctx context.Context, upload *model.Upload) (*repository.RepositoryResponse, error)
	UploadTest(c *gin.Context, upload *model.Upload) (*repository.RepositoryResponse, error)
	Read(read *model.Chunk) (repository.RepositoryResponse, error)
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
		File:     upload.File,
	})

	return RepositoryResponse, err
}

func (service *UploadService) Read(read *model.Chunk) (repository.RepositoryResponse, error) {
	resp, err := service.repository.Read(read)
	if err != nil {
		panic(err)
	}

	file, err := resp.File.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// readerToStdout(file, read.Offset, read.Limit)

	response := repository.RepositoryResponse{
		Success: true,
		File:    readerToStdout(file, read.Offset, read.Limit),
	}

	return response, err
}

func readerToStdout(r multipart.File, offset int64, limit int64) []byte {
	fmt.Printf("bytes=%d-%d", offset, limit)
	buf := make([]byte, limit)
	n, err := r.ReadAt(buf, offset)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
	return (buf[:n])
}
