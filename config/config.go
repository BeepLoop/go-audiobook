package config

type Config struct {
	IP          string
	Port        string
	ListenAddr  string
	VoiceId     string
	VoiceEngine string
	User        string
	ApiKey      string
	Dsn         string
}

type Setting struct {
	Id          string `json:"id"`
	VoiceEngine string `json:"voice_engine"`
	User        string `json:"playht_user"`
	ApiKey      string `json:"api_key"`
	Dsn         string `json:"dsn"`
}
