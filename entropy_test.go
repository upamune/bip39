package bip39

import (
	"fmt"
	"testing"
)

func TestEntropy_ToMnemonic(t *testing.T) {
	entropy, err := NewEntropy("133755ff")
	if err != nil {
		t.Fatal(err)
	}
	mnemonic, err := entropy.ToMnemonic("english")
	if err != nil {
		t.Error(err)
	}
	expected := "basket rival lemon"
	if mnemonic.Words != expected {
		t.Error(fmt.Errorf("mnemonic.Words(%s) != expected(%s", mnemonic.Words, expected))
	}
}
