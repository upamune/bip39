package bip39

import (
	"encoding/hex"
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

func TestMnemonic_IsValid(t *testing.T) {
	words := "imitate robot frame trophy nuclear regret saddle around inflict case oil spice"

	mnemonic := Mnemonic{Words: words}

	wl, _ := GetWordlists("english")
	if mnemonic.IsValid(wl) != true {
		t.Errorf("mnemonic.IsValid(wl) == false")
	}
}

func TestMnemonic_ToSeed(t *testing.T) {
	mnemonic := Mnemonic{Words: "basket actual"}
	seed := mnemonic.ToSeed("")
	seedHex := mnemonic.ToSeedHex("")

	if hex.EncodeToString(seed) != seedHex {
		t.Errorf("hex.EncodeToString(seed)(%s) != seedHex(%s)", hex.EncodeToString(seed), seedHex)
	}

	expectedHex := "5cf2d4a8b0355e90295bdfc565a022a409af063d5365bb57bf74d9528f494bfa4400f53d8349b80fdae44082d7f9541e1dba2b003bcfec9d0d53781ca676651f"
	if seedHex != expectedHex {
		t.Errorf("seedHex(%s) != expectedHex(%s)", seedHex, expectedHex)
	}

	wl, _ := GetWordlists("english")

	if mnemonic.IsValid(wl) != false {
		t.Error("mnemonic.IsValid(wl) == true")
	}
}
