package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"avenue/app/model"
	"avenue/app/repository"
	"avenue/app/service"
)

type UploadController struct {
	service service.IService
}

func Execute(service service.IService) *UploadController {
	controller := &UploadController{
		service: service,
	}

	return controller
}

func (controller *UploadController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Sprintln("get form err: ", file)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// headers, err := os.Open(file.Filename)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 	return
	// }

	// id := uuid.New().String()

	domain := &model.Upload{
		File:     file,
		FileName: file.Filename,
	}

	reponse, err := controller.service.Upload(c.Request.Context(), domain)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Path", reponse.Id)
	c.Status(http.StatusCreated)

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func (controller *UploadController) UploadTest(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	headers, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer headers.Close()

	upload := &model.Upload{
		FileName: file.Filename,
		File:     file,
	}

	reponse, err := controller.service.UploadTest(c, upload)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	fmt.Println("File uploaded successfully: ", file.Filename)
	c.Header("Path", reponse.Id)
	c.JSON(
		http.StatusOK,
		&repository.RepositoryResponse{Success: reponse.Success, Id: reponse.Id},
	)
}

func (controller *UploadController) ReadTest(c *gin.Context) {
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

	reponse, err := controller.service.Read(chunk)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println("controller.service.Read: ", err)
		return
	}

	fmt.Println("File read successfully: ", reponse)
	c.JSON(
		http.StatusOK,
		reponse,
	)
}

func (controller *UploadController) Routes(engine *gin.Engine) {
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/test-upload", controller.UploadTest)
	engine.GET("/test-upload", controller.ReadTest)

	upload := engine.Group("upload")
	upload.POST("/", controller.Upload)
}
