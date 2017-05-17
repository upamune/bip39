package bip39

import "fmt"

type Mnemonic struct {
	Words string
}

func (m Mnemonic) ToSeed(password string) []byte {
	return []byte{}
}

func (m Mnemonic) ToSeedHex(password string) string {
	return ""
}

func (m Mnemonic) ToEntropy(wordlist Wordlist) string {
	return ""
}

var defaultStrength = 128

// TODO: 名前変更
const eightBits = 8

var defaultRandomGenerator = RandomGeneratorImpl{}

func NewMnemonic(strength *int, generator RandomGenerator, wordlistLang string) (Mnemonic,error) {
	if strength == nil {
		strength = &defaultStrength
	}
	if generator == nil {
		generator = defaultRandomGenerator
	}

	r := generator.Generate(*strength / eightBits)
	hex := intToHex(r)
	e := Entropy{Hex: hex}

	return e.ToMnemonic(wordlistLang)
}

type RandomGenerator interface {
	Generate(seed int) int
}

type RandomGeneratorImpl struct {
}

func (rg RandomGeneratorImpl) Generate(seed int) int {
	return 0
}

func intToHex(i int) string {
	return fmt.Sprintf("%x", i)
}
