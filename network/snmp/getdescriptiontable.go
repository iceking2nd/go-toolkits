package snmp

import (
	"github.com/soniah/gosnmp"
	"strings"
	"time"
)

const OID_PREFIX_DESCRIPTION = ".1.3.6.1.2.1.31.1.1.1.18"
const OID_PREFIX_IFNAME = ".1.3.6.1.2.1.31.1.1.1.1"

func GetInterfaceDescriptionTable(ip string, port uint16, community string) (map[string]string, error) {
	g := &gosnmp.GoSNMP{
		Target:             ip,
		Port:               port,
		Transport:          "udp",
		Community:          community,
		Version:            gosnmp.Version2c,
		Timeout:            time.Duration(2) * time.Second,
		Retries:            3,
		ExponentialTimeout: true,
		MaxOids:            gosnmp.MaxOids,
	}
	if errConnect := g.Connect(); errConnect != nil {
		return nil, errConnect
	} else {
		defer g.Conn.Close()
		var interfaceTable []gosnmp.SnmpPDU
		var descriptionTable []gosnmp.SnmpPDU
		var errWalkInterface, errWalkDescription error
		interfaceTable, errWalkInterface = g.BulkWalkAll(OID_PREFIX_IFNAME)
		if errWalkInterface != nil {
			return nil, errWalkInterface
		}
		descriptionTable, errWalkDescription = g.BulkWalkAll(OID_PREFIX_DESCRIPTION)
		if errWalkDescription != nil {
			return nil, errWalkDescription
		}

		var inte = map[string]string{}
		var desc = map[string]string{}
		for _, v := range interfaceTable {
			inte[strings.TrimPrefix(v.Name, OID_PREFIX_IFNAME)] = string(v.Value.([]byte))
		}
		for _, v := range descriptionTable {
			desc[strings.TrimPrefix(v.Name, OID_PREFIX_DESCRIPTION)] = string(v.Value.([]byte))
		}
		var InterfaceDescriptionTable = map[string]string{}
		for k, v := range inte {
			if len(desc[k]) > 0 {
				InterfaceDescriptionTable[v] = desc[k]
			}
		}
		return InterfaceDescriptionTable, nil
	}
}
