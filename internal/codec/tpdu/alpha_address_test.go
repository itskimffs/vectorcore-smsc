package tpdu

import (
	"testing"

	"github.com/svinson1121/vectorcore-smsc/internal/codec"
)

// Address-Length of an alphanumeric address is the number of useful semi-octets (TS 23.040 9.1.2.5).
func TestAlphaAddressLength(t *testing.T) {
	for _, tc := range []struct {
		alpha       string
		semiOctets  int
		valueOctets int
	}{
		{"kim", 6, 3},
		{"KimTel", 11, 6},
		{"Kim1234", 13, 7},
		{"KimTel12", 14, 7},
		{"KimTelAS123", 20, 10},
	} {
		out, err := encodeAddress(codec.Address{Alpha: tc.alpha})
		if err != nil {
			t.Fatalf("%q: %v", tc.alpha, err)
		}
		if int(out[0]) != tc.semiOctets || len(out) != 2+tc.valueOctets || out[1] != 0xD0 {
			t.Errorf("%q: length %d, %d octets, TOA %#x; want %d, %d, 0xd0", tc.alpha, out[0], len(out)-2, out[1], tc.semiOctets, tc.valueOctets)
		}
		addr, n, err := decodeAddress(append(out, 0xFF), 0)
		if err != nil || n != len(out) || addr.Alpha != tc.alpha {
			t.Errorf("%q: decoded %q, consumed %d of %d, err %v", tc.alpha, addr.Alpha, n, len(out), err)
		}
	}
}
