package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/honpery-com/be-api/config"
)

func main() {
	router := gin.Default()

	router.Run(fmt.Sprintf(":%d", config.Port))
}
