package controller

import (
	"fmt"
	"net/http"

	"avenue/app/model"
	"avenue/app/repository"
	"avenue/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// headers, err := file.Open()
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 	return
	// }

	id := uuid.New().String()

	domain := &model.Upload{
		Headers:  file.Header,
		FileName: file.Filename,
		ReqId:    id,
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

	upload := &model.Upload{
		Headers:  file.Header,
		FileName: file.Filename,
	}

	reponse, err := controller.service.UploadTest(file, c, upload)

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

func (controller *UploadController) Routes(engine *gin.Engine) {
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/test-upload", controller.UploadTest)

	upload := engine.Group("upload")
	upload.POST("/", controller.Upload)
}
