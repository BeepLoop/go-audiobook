package types

import "github.com/BeepLoop/go-audiobook/config"

type PlayHTPayload struct {
	Text          string `json:"text"`
	Voice         string `json:"voice"`
	Quality       string `json:"quality"`
	Output_format string `json:"output_format"`
	Voice_engine  string `json:"voice_engine"`
}

func NewPayload(word string) *PlayHTPayload {
	payload := PlayHTPayload{
		Text:          word,
		Voice:         config.Env.VoiceId,
		Quality:       "high",
		Output_format: "mp3",
		Voice_engine:  config.Env.VoiceEngine,
	}

	return &payload
}
