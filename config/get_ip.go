package config

import (
	"errors"
	"net"
)

func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("failed to get local IP addr")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	return localAddr, nil
}
