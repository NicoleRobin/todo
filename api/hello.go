package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags exmaple
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/hello [get]
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}
