package api

import (
	"log"
	"sync"

	"github.com/BeepLoop/go-audiobook/types"
)

type ConvertResult struct {
	Result *types.Word
	Error  error
}

func CreateResult(word *types.Word, err error) *ConvertResult {
	return &ConvertResult{
		Result: word,
		Error:  err,
	}
}

func ConvertTTS(words []string) []*types.Word {
	var convertedWords []*types.Word

	wordsChan := make(chan *ConvertResult)
	var wg sync.WaitGroup

	log.Println("words: ", words)

	for _, word := range words {
		wg.Add(1)
		go ContactPlayHT(word, wordsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(wordsChan)
	}()

	for word := range wordsChan {
		if word.Error != nil {
			log.Println(word.Error.Error())
			continue
		}

		convertedWords = append(convertedWords, word.Result)
	}

	return convertedWords
}
