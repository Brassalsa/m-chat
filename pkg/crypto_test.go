package pkg

import (
	"testing"
)

func TestCrypto(t *testing.T) {

	const str = "password"
	newStr := str

	t.Log("testing hashing...")
	if err := HashString(&newStr); err != nil {
		t.Error("error hashing ", err)
	}
	t.Log("comparing hash")
	if ok := CompareHash(newStr, str); !ok {
		t.Error("hash comparison failed, not working as intended")
	}
}
