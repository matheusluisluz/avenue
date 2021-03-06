package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"avenue/app/model"
	"avenue/app/repository"
)

type IService interface {
	UploadFs(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error)
	UploadMem(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error)
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

func (service *UploadService) UploadFs(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error) {
	id := uuid.New().String()
	dist := fmt.Sprintf("%s-%s", id, upload.FileName)

	if err := c.SaveUploadedFile(&upload.Header, dist); err != nil {
		panic(err)
	}
	fileId, err := service.repository.Upload(&model.Upload{
		Path:     dist,
		FileName: upload.FileName,
		ID:       id,
		File:     upload.File,
		Header:   upload.Header,
	})

	response := &model.UploadResponse{
		Success: true,
		Id:      fileId,
	}

	return response, err
}

func (service *UploadService) UploadMem(c *gin.Context, upload *model.Upload) (*model.UploadResponse, error) {
	id := uuid.New().String()
	dist := fmt.Sprintf("%s-%s", id, upload.FileName)

	fileId, nil := service.repository.Upload(&model.Upload{
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
	file, err := service.repository.Read(read)
	if err != nil {
		panic(err)
	}

	fmt.Println("UploadService: ")
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

	fileSufix := fmt.Sprint(read.UploadID, "-nasdaq_symbols.csv")
	tmpfile, err := ioutil.TempFile("", fileSufix)
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(file); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}

	response := &model.ReadResponse{
		Success: true,
		File:    readerToStdout(tmpfile, read.Offset, read.Limit),
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
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
