package utils

import (
	"log"
	"sync"

	"github.com/audiobook/types"
)

type DownloadResult struct {
	Result *types.Word
	Error  error
}

func CreateResult(word *types.Word, err error) *DownloadResult {
	return &DownloadResult{
		Result: word,
		Error:  err,
	}
}

func MultithreadDownload(words []*types.Word) []*types.Word {
	downloadChan := make(chan *DownloadResult)
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go DownloadFile(word, downloadChan, &wg)
	}

	go func() {
		wg.Wait()
		close(downloadChan)
	}()

	var downloaded []*types.Word
	for word := range downloadChan {
		if word.Error != nil {
			log.Println(word.Error.Error())
			continue
		}

		downloaded = append(downloaded, word.Result)
	}

	return downloaded
}
