package server

import (
	"log"
	"net/http"

	"github.com/BeepLoop/go-audiobook/middleware"
	"github.com/BeepLoop/go-audiobook/store"
	"github.com/BeepLoop/go-audiobook/utils"
	"github.com/gin-gonic/gin"
)

func HandleStoryRoute(story *gin.RouterGroup) {
	story.GET("/", func(c *gin.Context) {
		html := utils.HtmlParser("stories.tmpl", "components/header.tmpl", "components/topbar.tmpl")
		Router.SetHTMLTemplate(html)

		log.Println("stories: ", store.Stories.Stories)

		c.HTML(http.StatusOK, "stories.tmpl", gin.H{
			"stories": store.Stories.Stories,
		})
	})

	story.GET("/new", middleware.CookieMonster, func(c *gin.Context) {
		html := utils.HtmlParser("newStory.tmpl", "components/header.tmpl", "components/adminTopbar.tmpl")
		Router.SetHTMLTemplate(html)

		c.HTML(http.StatusOK, "newStory.tmpl", nil)
	})

	story.GET("/:id", func(c *gin.Context) {
		html := utils.HtmlParser("readStory.tmpl", "components/header.tmpl", "components/topbar.tmpl")
		Router.SetHTMLTemplate(html)

		id := c.Params.ByName("id")

		story := store.GetStoryData(id)

		c.HTML(http.StatusOK, "readStory.tmpl", gin.H{
			"story": story,
		})
	})

	story.POST("/delete/:id", middleware.CookieMonster, func(c *gin.Context) {
		id := c.Params.ByName("id")

		err := store.DeleteStoryLocal(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}
