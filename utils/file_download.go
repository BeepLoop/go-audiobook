package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/audiobook/types"
)

func DownloadFile(word *types.Word, ch chan<- *DownloadResult, wg *sync.WaitGroup) *DownloadResult {
	defer wg.Done()

	fullFilePath := fmt.Sprintf("assets/downloads/%v.mp3", word.Word)
	file, err := os.Create(fullFilePath)
	if err != nil {
		log.Println(err)
		result := CreateResult(nil, err)
		ch <- result
		return result
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(word.Location)
	if err != nil {
		result := CreateResult(nil, err)
		ch <- result
		return result
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		result := CreateResult(nil, err)
		ch <- result
		return result
	}
	log.Printf("downloaded audio for: %s", word.Word)

	downloaded := types.NewWordData(word.Word, fmt.Sprintf("/downloads/%v.mp3", word.Word))
	result := CreateResult(downloaded, nil)

	ch <- result

	return result
}
