package controller

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"avenue/app/config"
	"avenue/app/model"
	"avenue/app/service"
)

type UploadController struct {
	service service.IService
	config  config.Configuration
}

func Execute(service service.IService, config config.Configuration) *UploadController {
	controller := &UploadController{
		service: service,
		config:  config,
	}

	return controller
}

func (controller *UploadController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	reponse, err := controller.fsOrMemory(file, c)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	fmt.Println("File uploaded successfully: ", file.Filename)
	c.JSON(
		http.StatusOK,
		reponse,
	)
}

func (controller *UploadController) Read(c *gin.Context) {
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	chunk := &model.Chunk{
		UploadID: c.Query("id"),
		Offset:   offset,
		Limit:    limit,
	}
	fmt.Println("chunk: ", chunk)
	if err := c.BindQuery(chunk); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	response, err := controller.service.Read(chunk)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println("controller.service.Read: ", err)
		return
	}

	fmt.Println("File read successfully: ", response.Success)
	c.JSON(
		http.StatusOK,
		response,
	)
}

func (controller *UploadController) Routes(engine *gin.Engine) {
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", controller.Upload)
	engine.GET("/upload", controller.Read)
}

func (controller *UploadController) UploadMemory(file *multipart.FileHeader, c *gin.Context) (*model.UploadResponse, error) {
	headers, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return nil, err
	}
	defer headers.Close()

	upload := &model.Upload{
		FileName: file.Filename,
		File:     headers,
	}

	return controller.service.UploadMem(c, upload)
}

func (controller *UploadController) UploadFs(file *multipart.FileHeader, c *gin.Context) (*model.UploadResponse, error) {
	headers, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return nil, err
	}
	defer headers.Close()

	upload := &model.Upload{
		FileName: file.Filename,
		Header:   *file,
		File:     headers,
	}

	return controller.service.UploadFs(c, upload)
}

func (controller *UploadController) fsOrMemory(file *multipart.FileHeader, c *gin.Context) (*model.UploadResponse, error) {
	switch controller.config.Upload {
	case "fs":
		return controller.UploadFs(file, c)
	case "memory":
		return controller.UploadMemory(file, c)
	}
	return nil, errors.New("backend not identified")
}
