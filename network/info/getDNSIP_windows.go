//+build windows

package info

import (
	"fmt"
	"github.com/kbinani/win"
	"net"
	"unsafe"
)

func GetDNSIP() ([]net.IP, error) {
	dns := []net.IP{}
	info := win.FIXED_INFO_W2KSP1{}
	size := uint32(unsafe.Sizeof(info))
	r := win.GetNetworkParams(&info, &size)
	if r == 0 {
		for ai := &info.DnsServerList; ai != nil; ai = ai.Next {
			d := fmt.Sprintf("%v.%v.%v.%v", ai.Context&0xFF, (ai.Context>>8)&0xFF, (ai.Context>>16)&0xFF, (ai.Context>>24)&0xFF)
			dns = append(dns, net.ParseIP(d))
		}
	} else if r == win.ValueOverflow {
		newBuffers := make([]byte, size)
		netParams := (win.PFIXED_INFO)(unsafe.Pointer(&newBuffers[0]))
		win.GetNetworkParams(netParams, &size)
		for ai := &netParams.DnsServerList; ai != nil; ai = ai.Next {
			d := fmt.Sprintf("%v.%v.%v.%v", ai.Context&0xFF, (ai.Context>>8)&0xFF, (ai.Context>>16)&0xFF, (ai.Context>>24)&0xFF)
			dns = append(dns, net.ParseIP(d))
		}
	}
	return dns, nil

}
