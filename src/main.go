package main

import (
	"avenue/app/controller"
	"avenue/app/repository"
	"avenue/app/service"
	"fmt"
	"io/ioutil"

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

	controller.Routes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
