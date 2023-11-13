package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieMonster(c *gin.Context) {
	_, err := c.Cookie("admin")
	if err != nil {
		log.Println("error cookie: ", err)
		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/admin/login")
        c.Abort()
		return
	}

	c.Next()
}
