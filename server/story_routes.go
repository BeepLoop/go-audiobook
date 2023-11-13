package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/audiobook/middleware"
	"github.com/audiobook/store"
	"github.com/gin-gonic/gin"
)

type NewStory struct {
	Status string
}

func HandleStoryRoute(story *gin.RouterGroup) {
	story.GET("/", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/stories.html")

		stories, err := store.GetStories()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, "stories.html", gin.H{
			"stories": stories,
		})
	})

	story.GET("/new", middleware.CookieMonster, func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/newStory.html")

		c.HTML(http.StatusOK, "newStory.html", nil)
	})

	story.GET("/:id", func(c *gin.Context) {
		Router.LoadHTMLGlob("views/templates/readStory.html")
		paramsId := c.Params.ByName("id")

		id, err := strconv.Atoi(paramsId)
		if err != nil {
			c.HTML(http.StatusBadRequest, "readStory.html", nil)
			return
		}

		story, words, err := store.GetStoryData(id)
		if err != nil {
			c.HTML(http.StatusBadRequest, "readStory.html", nil)
			return
		}

		c.HTML(http.StatusOK, "readStory.html", gin.H{
			"story": story,
			"words": words,
		})
	})

	story.POST("/delete/:id", middleware.CookieMonster, func(c *gin.Context) {
		paramsId := c.Params.ByName("id")
		id, err := strconv.Atoi(paramsId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
			return
		}

		err = store.DeleteStory(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}
