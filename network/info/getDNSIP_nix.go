//+build !plan9,!windows

package info

import (
	"io/ioutil"
	"net"
	"strings"
)

func GetDNSIP() ([]net.IP, error) {
	if content, errReadFile := ioutil.ReadFile("/etc/resolv.conf"); errReadFile != nil {
		return nil, errReadFile
	} else {
		linesArray := strings.Split(string(content), "\n")
		var dns = []net.IP{}
		for _, v := range linesArray {
			if strings.Split(v, " ")[0] == "nameserver" {
				dns = append(dns, net.ParseIP(strings.Split(v, " ")[1]))
			}
		}
		return dns, nil
	}
}
