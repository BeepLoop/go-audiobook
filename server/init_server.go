package server

import (
	"encoding/gob"
	"net/http"

	"github.com/audiobook/middleware"
	"github.com/audiobook/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitServer() {
	gob.Register(types.Admin{})

	Router = gin.Default()

	sessionStore := InitSession()

	Router.Use(cors.Default())
	Router.Use(middleware.MimeSetter)
	Router.Use(sessions.Sessions("admin", sessionStore))

	Router.Static("/styles", "views/styles/")
	Router.Static("/scripts", "views/scripts/")
	Router.Static("/downloads", "assets/downloads/")
	Router.Static("/resource", "assets/public/")
	Router.Static("/audios", "assets/audios/")
	Router.Static("/thumbnails", "assets/thumbnails/")

	Router.GET("/", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/index.html")

		c.HTML(http.StatusOK, "index.html", nil)
	})

	Router.NoRoute(func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/404.html")

		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	story := Router.Group("/story")
	HandleStoryRoute(story)

	admin := Router.Group("/admin")
	HandleAdminRoute(admin)

	upload := Router.Group("/upload")
	HandleUploadRoute(upload)
}
