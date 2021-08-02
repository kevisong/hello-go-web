package server

import (
	"net/http"
	"strconv"

	"github.com/KEVISONG/hello-go-web/pkg/models"
	"github.com/gin-gonic/gin"
)

func userGet(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{Code: 1, Message: err.Error()})
		return
	}
	user, err := db.GetUserByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusOK, models.Response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Message: "success", Data: user})
}
