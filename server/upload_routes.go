package server

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/BeepLoop/go-audiobook/api"
	"github.com/BeepLoop/go-audiobook/store"
	"github.com/BeepLoop/go-audiobook/types"
	"github.com/BeepLoop/go-audiobook/utils"
	"github.com/gin-gonic/gin"
)

type StoryFiles struct {
	Type *multipart.FileHeader
	Dist string
}

func HandleUploadRoute(upload *gin.RouterGroup) {
	upload.POST("/story", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Println(err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/story/new?status=failed")
			return
		}

		audio := form.File["audio"][0]
		thumbnail := form.File["thumbnail"][0]
		utils.RenameFile(audio, form.Value["title"][0])
		utils.RenameFile(thumbnail, form.Value["title"][0])

		audioDist := "/audios/" + audio.Filename
		thumbnailDist := "/thumbnails/" + thumbnail.Filename

		words := strings.Split(form.Value["words"][0], ", ")
		convertedWords := api.ConvertTTS(words)

		if len(convertedWords) < 1 {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/story/new?status=failed")
			return
		}

		downloadedWords := utils.MultithreadDownload(convertedWords)
		fmt.Println("downloaded words")
		for _, word := range downloadedWords {
			fmt.Println("word: ", word)
		}

		if len(downloadedWords) < 1 {
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/story/new?status=failed")
			return
		}

		randId, err := utils.GenerateUUID()
		if err != nil {
			log.Println(err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/story/new?status=failed")
			return
		}

		story := types.Story{
			Id:      randId,
			Title:   form.Value["title"][0],
			Author:  form.Value["author"][0],
			Content: form.Value["story"][0],
			Audio:   audioDist,
			Image:   thumbnailDist,
			Words:   downloadedWords,
		}

		err = store.SaveStoryLocal(&story)
		if err != nil {
			log.Println(err)
			c.Request.Method = "GET"
			c.Redirect(http.StatusSeeOther, "/story/new?status=failed&reason="+err.Error())
			return
		}

		files := []StoryFiles{
			{
				Type: audio,
				Dist: "assets" + audioDist,
			},
			{
				Type: thumbnail,
				Dist: "assets" + thumbnailDist,
			},
		}

		for _, file := range files {
			err := c.SaveUploadedFile(file.Type, file.Dist)
			if err != nil {
				log.Printf("Failed to save: %v", file.Type.Filename)
				c.Request.Method = "GET"
				c.Redirect(http.StatusSeeOther, "/story/new?status=failed&reason="+err.Error())
				return
			}
		}

		c.Request.Method = "GET"
		c.Redirect(http.StatusSeeOther, "/story/new?status=success")
	})
}
