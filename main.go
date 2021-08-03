package main

import (
	"log"
	"telkomGin/config"
	"telkomGin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Routes(r)
	config.Connect()
	log.Fatal(r.Run(":8080"))
}
