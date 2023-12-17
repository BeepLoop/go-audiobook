package config

import (
	"strings"
)

var Env *Config

func Init(port string) error {
	voiceSetting := LoadSettings()

	ip, err := GetOutboundIP()
	if err != nil {
		return err
	}

	localAddr := strings.Split(ip, ":")[0]
	Env = &Config{
		IP:          localAddr,
		Port:        port,
		ListenAddr:  localAddr + ":" + port,
		VoiceId:     voiceSetting.Id,
		VoiceEngine: voiceSetting.VoiceEngine,
		User:        voiceSetting.User,
		ApiKey:      voiceSetting.ApiKey,
	}

	return nil
}
