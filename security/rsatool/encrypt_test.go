package rsatool

import (
	"reflect"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	var privatekey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAqq/X9ajhauK/fAGA7BCB1qZ0LiRTP1EDlf4SegcDpVkIspag
fq8sood1iwCUgSFKH1zgwZsneZ1RPEcts640W9jWPDRoSW2caLstpNT6RCs1fURV
czdityLQfNNby2s7R1XazS97U+p4t14PJKQhrtdJcJirQ5eso3BwvBv8Bx10Ow34
c4u44MrhBCVwAw9InrjyTtA/EttORnYVmZXazpcaVeC8OtAA5WtYyEhfonbA9UFw
8re/jSjNijMsKgyguBoDEFtgUrbmgsrpjPur+GZcvzMvlXvxK72O78nH7zXBHkni
8iQMFFHR2Vr/mgKAVmIRbLj08JKD49oxm2P1RwIDAQABAoIBACvrKRdVyAK3q2ud
+21CZhY0KErDIAR3tQOAJ4LskRyXhyxko5FVQ5fiYPVORefgB/F/9Xj7kjGxr/hh
nDnShjhn1+AP4BGcuZUdLIZwIqEtuQqtbYtC4gybUaf8vm7d/ZfCWJVZNVOHTl5o
Pl9AOEaDsKrfLZoryckYInnUfYnat1rJT7xiU2Bw5iQXUjlZq0T0N7sssoHki3oh
BKIuzTHZac77XuiL6c4xjLuQfPE3T/zV5v+kDsg9Gr+2wJuXyOuU/PcAVOM5Wd3p
qlJrLcqoF1wMd3qt+LBi4qklf/wdpG2O1e0Nb5ynIEttaNca4rZeYPO/PYET6EAg
D2ZTbhECgYEA44WxwlPKvv4w4ZkC3akS96Uci6UP0pw05YMApKrFIsxnEsl8r8AG
ARCBdc7SwpklZjqAoM5rye/pJPVzYA5Xw1YokxTL72taarjF8PP7j77IOYEQmZyz
5qsEDEg/2uxO3O1hcNGizkuSt6ZvfDRFFpbfP34nIcXygzQssMhSZQ8CgYEAwA0F
dVPuODj6a27BnrbGRvIuEbAqslk+SoryaoX7layOfezVg7LXM5zvbcoPe2ZAxgM2
7M+N+dSChD13DzgLNRvmCxLX/laXDRWFv9WBRLc6fP3cVaWgmEJWJh/wK0KMCf02
TXs41RAppQHQOokCiSONszfv8byOKOp0WnSwnEkCgYEA06uvC9Z+uh+sBEYhF21T
mxd+LiC7TrpKr/enoorJaq11e1H+cVeZfmYHuz1WcPqfg7MLtg800qMjaWH9G7Fc
TvuaNfPUIwLiFIWgiaHDOSFKD82WTwMtj30s/+w7lAXi/MDPtXEjIP/IGJl8ALZ0
TCbCvcQ0zcPl8dCZf0Ju/VsCgYEAi/KxMI9pv6Sdr131T//mutpOAeu7IV3xC4b3
IfLio24sI81Kxf6z3VOWEg41e/nGZ8T96FFTJ2PZfB4CPZ/cYuYim50uspNM3Wb6
oREncviGcYDYZ70bKvjkKEIL1KpYvAq94HZTssdlYW3R3GXs6Bx1bMztV93MRjqb
cTMpY6kCgYAiGsle9NVtkwex1ZrxZOlCU3NWrWwRNgFh9puBXuIVZ6WbvUIP6SqO
M/C5qXEBsdTIUkIDPFPmE5R9DII1yj0NngRskrAL+FWHAAtT71iX5JKu4xgG1C2v
yEC1h9uImI7kMITcgUJU1L9UmoF6qp58+ARj2r5lZmqDkKRDDpkSxQ==
-----END RSA PRIVATE KEY-----`)
	var publickey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqq/X9ajhauK/fAGA7BCB
1qZ0LiRTP1EDlf4SegcDpVkIspagfq8sood1iwCUgSFKH1zgwZsneZ1RPEcts640
W9jWPDRoSW2caLstpNT6RCs1fURVczdityLQfNNby2s7R1XazS97U+p4t14PJKQh
rtdJcJirQ5eso3BwvBv8Bx10Ow34c4u44MrhBCVwAw9InrjyTtA/EttORnYVmZXa
zpcaVeC8OtAA5WtYyEhfonbA9UFw8re/jSjNijMsKgyguBoDEFtgUrbmgsrpjPur
+GZcvzMvlXvxK72O78nH7zXBHkni8iQMFFHR2Vr/mgKAVmIRbLj08JKD49oxm2P1
RwIDAQAB
-----END PUBLIC KEY-----`)
	var data = []byte("iceking2nd test")
	if p, err := RsaEncrypt(data, publickey); err != nil {
		t.Error(err)
	} else {
		c, _ := RsaDecrypt(p, privatekey)
		if !reflect.DeepEqual(c, data) {
			t.Errorf("RsaEncrypt(%v,%s) = %s, expected %s", p, publickey, c, data)
		}
	}
}
