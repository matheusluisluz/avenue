package main

import (
	"avenue/app/controller"
	"avenue/app/repository"
	"avenue/app/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func UploadTest(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {
	fmt.Println("Hello World")

	// router.GET("/", HomePage)
	// router.POST("/", PostHomePage)
	// router.GET("/query", QueryString)
	// router.GET("/path/:name/:age", PathParams)

	repository := repository.Execute()
	service := service.Execute(repository)
	controller := controller.Execute(service)

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/test-upload", UploadTest)

	controller.Routes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
