package bip39

import (
	"fmt"
	"testing"
)

func TestEntropy_ToMnemonic(t *testing.T) {
	entropy := Entropy{Hex: "133755ff"}
	mnemonic, err := entropy.ToMnemonic("english")
	if err != nil {
		t.Error(err)
	}
	expected := "basket rival lemon"
	if mnemonic.Words != expected {
		t.Error(fmt.Errorf("mnemonic.Words(%s) != expected(%s", mnemonic.Words, expected))
	}
}
