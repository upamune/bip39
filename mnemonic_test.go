package bip39

import (
	"fmt"
	"testing"
)

type mockGenerator struct{}

func (g mockGenerator) Generate(size int) []byte {
	seed := "qwertyuiopasdfghjklzxcvbnm[];,./"
	b := []byte(seed)[:size]

	return b
}

func TestNewMnemonic(t *testing.T) {
	mnemonic, err := NewMnemonic(nil, mockGenerator{}, "english")
	if err != nil {
		t.Fatal(err)
	}

	expected := "imitate robot frame trophy nuclear regret saddle around inflict case oil spice"
	if mnemonic.Words != expected {
		t.Fatal(fmt.Errorf("mnemonic.Words(%s) != expected(%s)", mnemonic.Words, expected))
	}
}

func TestMnemonic_ToEntropy(t *testing.T) {
	mnemonic := Mnemonic{Words: "basket rival lemon"}
	wl, _ := GetWordlists("english")
	entropy, err := mnemonic.ToEntropy(wl)
	if err != nil {
		t.Fatal(err)
	}

	expected := "133755ff"

	if expected != entropy {
		t.Fatal(fmt.Sprintf(`expected("%s") != entropy("%s")`, expected, entropy))
	}
}
