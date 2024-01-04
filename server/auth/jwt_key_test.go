package auth

import (
	"bytes"
	"testing"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(RawPublicKey, want) {
		t.Errorf("want %s, but got %s", want, RawPublicKey)
	}
}
