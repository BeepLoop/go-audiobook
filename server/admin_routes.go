package server

import (
	"log"
	"net/http"
	"reflect"

	"github.com/audiobook/config"
	"github.com/audiobook/middleware"
	"github.com/audiobook/store"
	"github.com/audiobook/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleAdminRoute(admin *gin.RouterGroup) {
	admin.GET("/stories", middleware.CookieMonster, func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/adminStories.html")

		stories, err := store.GetStories()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, "adminStories.html", gin.H{
			"stories": stories,
		})
	})

	admin.GET("/login", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/login.html")
		c.HTML(http.StatusOK, "login.html", nil)
	})

	admin.POST("/login", func(c *gin.Context) {
		var credentials types.User

		err := c.ShouldBindWith(&credentials, binding.Form)
		if err != nil {
			log.Println(err)
			return
		}

		user, err := store.GetCredentials(credentials.Username)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/admin/login?status=failed")
			return
		}

		if !reflect.DeepEqual(*user, credentials) {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/admin/login?status=failed")
			return
		}

		session := sessions.Default(c)
		session.Set("admin", *user)
		session.Save()

		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/stories")
	})

	admin.POST("/logout", middleware.CookieMonster, func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		cookie, _ := c.Cookie("admin")
		c.SetCookie("admin", cookie, -1, "/", config.Env.IP, false, true)

		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/")
	})
}
