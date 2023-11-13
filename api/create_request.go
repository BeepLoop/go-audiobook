package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/audiobook/config"
	"github.com/audiobook/types"
)

func CreateRequest(payload *types.PlayHTPayload) (*http.Request, error) {
	json_payload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", playHTUrl, bytes.NewBuffer(json_payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "text/event-stream")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("AUTHORIZATION", config.Env.ApiKey)
	req.Header.Add("X-USER-ID", config.Env.User)

	return req, nil
}
