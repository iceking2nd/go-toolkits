package info

import (
	"net"
	"strings"
)

func GetOutBoundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return net.ParseIP(localAddr[0:idx]), nil
}
