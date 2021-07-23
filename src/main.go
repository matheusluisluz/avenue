package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"

	"avenue/app/config"
	"avenue/app/controller"
	"avenue/app/repository"
	"avenue/app/service"
)

func main() {
	fmt.Println("Hello World")

	configuration := config.Execute()
	fmt.Println("configuration", configuration.Server.Port)
	fmt.Println("configuration", configuration.Upload)
	cacheConfig := bigcache.DefaultConfig(10 * time.Minute)

	repository := repository.Execute(cacheConfig)
	service := service.Execute(repository)
	controller := controller.Execute(service, configuration)

	router := gin.Default()

	controller.Routes(router)

	if err := router.Run(configuration.Server.Port); err != nil {
		panic(err)
	}
}
