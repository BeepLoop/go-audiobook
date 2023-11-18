package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/BeepLoop/go-audiobook/types"
	"github.com/BeepLoop/go-audiobook/utils"
)

var playHTUrl = "https://api.play.ht/api/v2/tts"

func ContactPlayHT(word string, ch chan<- *ConvertResult, wg *sync.WaitGroup) *ConvertResult {
	client := &http.Client{Timeout: time.Second * 30}
	payload := types.NewPayload(word)
	var result *ConvertResult

	defer wg.Done()

	req, err := CreateRequest(payload)
	if err != nil {
		result = CreateResult(nil, err)
		ch <- result
		return result
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request for word: %v responded with error: %v", word, err)
		result = CreateResult(nil, err)
		ch <- result
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Server responded with: %v", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		log.Println("body: ", string(body))
		result = CreateResult(nil, errors.New(errorMessage))
		ch <- result
		return result
	}

	eventArray, err := EventConsumer(resp.Body)
	if err != nil {
		result = CreateResult(nil, err)
		ch <- result
		return result
	}

	eventMap := utils.ArrayToMap(eventArray)
	if eventMap["stage"] != "complete" {
		errorMessage := fmt.Sprintf("Failed to convert to audio :%v", word)
		result = CreateResult(nil, errors.New(errorMessage))
		ch <- result
		return result
	}
	convertedWord := types.NewWordData(word, eventMap["url"])
	result = CreateResult(convertedWord, nil)

	ch <- result

	return result
}
