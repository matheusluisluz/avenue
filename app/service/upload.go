package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"avenue/app/model"
	"avenue/app/repository"
)

type IService interface {
	Upload(ctx context.Context, upload *model.Upload) (*model.UploadResponse, error)
	UploadTest(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error)
	Read(read *model.Chunk) (*model.ReadResponse, error)
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

func (service *UploadService) Upload(ctx context.Context, upload *model.Upload) (*model.UploadResponse, error) {
	id := uuid.New().String()
	location := fmt.Sprintf("%s-%s", id, upload.FileName)
	fileId, err := service.repository.Upload(ctx, &model.Upload{
		Path:     location,
		FileName: upload.FileName,
		ID:       id,
	})

	response := &model.UploadResponse{
		Success: true,
		Id:      fileId,
	}

	return response, err
}

func (service *UploadService) UploadTest(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error) {
	id := uuid.New().String()
	dist := fmt.Sprintf("%s-%s", id, upload.FileName)
	// if err := c.SaveUploadedFile(file, dist); err != nil {
	// 	panic(err)
	// }

	fileId, nil := service.repository.Upload(c.Request.Context(), &model.Upload{
		Path:     dist,
		FileName: upload.FileName,
		ID:       id,
		File:     upload.File,
	})

	response := &model.UploadResponse{
		Success: true,
		Id:      fileId,
	}

	return response, nil
}

func (service *UploadService) Read(read *model.Chunk) (*model.ReadResponse, error) {
	resp, err := service.repository.Read(read)
	if err != nil {
		panic(err)
	}

	fmt.Println("UploadService: ", resp)
	// file, err := resp.File.Read()
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// headers, err := resp.File.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// defer headers.Close()

	response := &model.ReadResponse{
		Success: true,
		File:    resp,
	}

	return response, err
}

func readerToStdout(r multipart.File, offset int64, limit int64) []byte {
	fmt.Println("offset: ", offset)
	fmt.Println("limit: ", limit)
	buf := make([]byte, limit)
	n, err := r.ReadAt(buf, offset)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
	return (buf[:n])
}
