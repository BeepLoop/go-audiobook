package server

import (
	"encoding/gob"
	"net/http"

	"github.com/BeepLoop/go-audiobook/middleware"
	"github.com/BeepLoop/go-audiobook/types"
	"github.com/BeepLoop/go-audiobook/utils"
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
		html := utils.HtmlParser("index.html", "components/header.html")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusOK, "index.html", nil)
	})

	Router.NoRoute(func(c *gin.Context) {
		html := utils.HtmlParser("404.html", "components/header.html")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	story := Router.Group("/story")
	HandleStoryRoute(story)

	admin := Router.Group("/admin")
	HandleAdminRoute(admin)

	upload := Router.Group("/upload")
	HandleUploadRoute(upload)
}
