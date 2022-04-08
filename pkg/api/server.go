package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/golang-base/pkg/configs"
)

const statusPath = "/status"

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	r := gin.Default()
	r.GET(statusPath, statusHandler)

	if err := r.Run(fmt.Sprintf(":%d", configs.GetPort())); err != nil {
		panic(err)
	}
}