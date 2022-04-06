package controller

import (
	"net/http"
	"strconv"

	"github.com/KEVISONG/hello-go-web/pkg/database"
	"github.com/KEVISONG/hello-go-web/pkg/models"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
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

func UserPost(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{Code: 1, Message: err.Error()})
		return
	}
	id, err := db.PostUser(user)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{Code: 1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Message: "success", Data: id})
}

var db *database.Store

// Init initialize the api server
func Init(d *database.Store) {
	db = d
}
