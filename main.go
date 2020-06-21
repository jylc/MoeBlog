package main

import (
	_ "MoeBlog/docs"
	"MoeBlog/handle"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	handle.InitRoutes(r)
	fmt.Println("[main] listen on 8080 port")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("[main] start error!!!")
	}
}
