package rsatool

import (
	"encoding/pem"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	var size = 2048
	if prik, pubk, err := GenRsaKey(size); err != nil {
		t.Error(err)
	} else {
		pubblock, _ := pem.Decode(pubk)
		if pubblock == nil {
			t.Error("public key error")
		}
		priblock, _ := pem.Decode(prik)
		if priblock == nil {
			t.Error("private key error")
		}
	}
}
