package server

import (
	"fmt"
	"net/http"

	"github.com/KEVISONG/hello-go-web/pkg/config/server"
	"github.com/KEVISONG/hello-go-web/pkg/database"
	"github.com/gin-gonic/gin"
)

func route(e *gin.Engine) {

	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/api").Group("/v1")
	v1.GET("/user/:id", userGet)

}

// Run runs server
func Run(c server.Config) {

	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	route(engine)

	engine.Run(fmt.Sprintf(":%d", c.Port))

}

var db *database.Store

func Init(d *database.Store) {
	db = d
}
