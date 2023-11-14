package server

import (
	"log"
	"net/http"

	"github.com/audiobook/middleware"
	"github.com/audiobook/store"
	"github.com/gin-gonic/gin"
)

func HandleStoryRoute(story *gin.RouterGroup) {
	story.GET("/", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/stories.html")

		log.Println("stories: ", store.Stories.Stories)

		c.HTML(http.StatusOK, "stories.html", gin.H{
			"stories": store.Stories.Stories,
		})
	})

	story.GET("/new", middleware.CookieMonster, func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/newStory.html")

		c.HTML(http.StatusOK, "newStory.html", nil)
	})

	story.GET("/:id", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/readStory.html")
		id := c.Params.ByName("id")

		story := store.GetStoryData(id)

		c.HTML(http.StatusOK, "readStory.html", gin.H{
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
