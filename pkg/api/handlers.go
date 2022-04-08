package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusHandler TODO
// @Summary Shows the status of the application.
// @Description TODO
// @Accept json
// @Produce json
// @Router /status [get]
func StatusHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}