package config

import "fmt"

var Env *Config

func Init(port string) error {
	voiceSetting := LoadSettings()

	ip, err := GetOutboundIP()
	if err != nil {
		return err
	}

	Env = &Config{
		IP:          fmt.Sprintf("%v", ip),
		Port:        port,
		ListenAddr:  fmt.Sprintf("%v:%v", ip, port),
		VoiceId:     voiceSetting.Id,
		VoiceEngine: voiceSetting.VoiceEngine,
		User:        voiceSetting.User,
		ApiKey:      voiceSetting.ApiKey,
	}

	return nil
}
