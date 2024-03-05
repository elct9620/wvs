package session_test

import (
	"testing"

	"github.com/elct9620/wvs/pkg/session"
)

func TestEncrypt(t *testing.T) {
	key := "1234567890123456"
	data := "b285e7fb-2ed6-4d4c-a013-c0c0655f5769"

	encrypted, err := session.Encrypt([]byte(data), []byte(key))
	if err != nil {
		t.Error(err)
	}

	decrypted, err := session.Decrypt(encrypted, []byte(key))
	if err != nil {
		t.Error(err)
	}

	if string(decrypted) != data {
		t.Errorf("expected %s, got %s", data, decrypted)
	}
}
