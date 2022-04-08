package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func statusHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}