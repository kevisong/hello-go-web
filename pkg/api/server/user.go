package server

import (
	"net/http"

	"github.com/KEVISONG/hello-go-web/pkg/models"
	"github.com/gin-gonic/gin"
)

func userGet(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, models.Response{Code: 0, Message: "success", Data: id})
}
