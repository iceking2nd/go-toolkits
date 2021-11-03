package ipv4

import (
	"encoding/binary"
	"net"
	"strconv"
)

func Uint64toIP4(ipInt uint64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatUint((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatUint((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatUint((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatUint(ipInt&0xff, 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func IPv4toUint64(ipStr string) uint64 {
	return uint64(binary.BigEndian.Uint32(net.ParseIP(ipStr)[12:16]))
}
