package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"githup.com/dierbei/go-helm-api/config"
	"githup.com/dierbei/go-helm-api/internal/router"
	"log"
)

func main() {
	c := config.GetConfig().Application

	engine := gin.Default()
	if c.Mode != "debug" {
		engine = gin.New()
		engine.Use(gin.Recovery())
	}

	router.Router(engine)

	if err := engine.Run(fmt.Sprintf(":%d", c.Port)); err != nil {
		log.Fatal(err)
	}
}
