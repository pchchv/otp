package otp_test

import (
	"testing"

	"github.com/pchchv/otp"
)

func TestITob(t *testing.T) {
	var i int64 = 1524486261
	expect := []byte{0, 0, 0, 0, 90, 221, 208, 117}

	if string(expect) != string(otp.Itob(i)) {
		t.Error("ITob error")
	}
}

func TestBuildUri(t *testing.T) {
	s := otp.BuildUri(
		"totp",
		"4S62BZNFXXSZLCRO",
		"xlzd",
		"SomeOrg",
		"sha1",
		0,
		6,
		0,
	)
	expected := "otpauth://totp/SomeOrg:xlzd?issuer=SomeOrg&secret=4S62BZNFXXSZLCRO"
	if s != expected {
		t.Errorf("BuildUri test failed.\n\texpected: %s,\n\tactual: %s", expected, s)
	}

	s2 := otp.BuildUri(
		"hotp",
		"XXSZLCRO4S62BZNF",
		"mergenchik@gmail.com",
		"github.com",
		"sha1",
		0,
		6,
		0)
	expected2 := "otpauth://hotp/github.com:mergenchik@gmail.com?counter=0&issuer=github.com&secret=XXSZLCRO4S62BZNF"
	if s2 != expected2 {
		t.Errorf("BuildUri test failed.\n\texpected: %s,\n\tactual: %s", expected2, s2)
	}
}
