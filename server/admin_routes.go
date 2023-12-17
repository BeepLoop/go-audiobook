package server

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/BeepLoop/go-audiobook/config"
	"github.com/BeepLoop/go-audiobook/middleware"
	"github.com/BeepLoop/go-audiobook/store"
	"github.com/BeepLoop/go-audiobook/types"
	"github.com/BeepLoop/go-audiobook/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleAdminRoute(admin *gin.RouterGroup) {
	admin.GET("/stories", middleware.CookieMonster, func(c *gin.Context) {
		html := utils.HtmlParser("adminStories.tmpl", "components/header.tmpl", "components/adminTopbar.tmpl")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusOK, "adminStories.tmpl", gin.H{
			"stories": store.Stories.Stories,
		})
	})

	admin.GET("/login", func(c *gin.Context) {
		html := utils.HtmlParser("login.tmpl", "components/header.tmpl", "components/topbar.tmpl")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	admin.POST("/login", func(c *gin.Context) {
		var credentials types.Admin

		err := c.ShouldBindWith(&credentials, binding.Form)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/admin/login/?status=failed")
			return
		}

		if utils.IsInputEmpty(&credentials) {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/admin/login/?status=failed")
		}

		if !reflect.DeepEqual(store.Credential, credentials) {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/admin/login?status=failed")
			return
		}

		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, "Login Success")
		fmt.Println(colored)

		session := sessions.Default(c)
		session.Set("admin", credentials)
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

	admin.GET("/settings", func(c *gin.Context) {
		html := utils.HtmlParser("settings.tmpl", "components/header.tmpl", "components/adminTopbar.tmpl")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusOK, "settings.tmpl", nil)
	})
}
