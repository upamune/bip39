package bip39

import (
	"testing"
	"github.com/k0kubun/pp"
)

func TestEntropy_ToMnemonic(t *testing.T) {
	entropy := Entropy{Hex: "133755ff"}
	mnemonic, err := entropy.ToMnemonic("english")
	if err != nil {
		t.Error(err)
	}
	pp.Println(mnemonic)
}
