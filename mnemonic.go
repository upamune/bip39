package bip39

import (
	"bytes"
	"encoding/hex"
	"math"
	"strconv"
	"strings"

	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/text/unicode/norm"
)

type Mnemonic struct {
	Words string
}

func (m Mnemonic) ToSeed(password string) []byte {
	return []byte{}
}

func (m Mnemonic) ToSeedHex(password string) string {
	return ""
}

func (m Mnemonic) ToEntropy(wordlist Wordlist) (string, error) {
	words := strings.Split(normalizeString(m.Words), " ")
	if len(words)%3 != 0 {
		return "", errors.New("Invalid mnemonic")
	}

	indexes := []int{}
	for _, word := range words {
		idx, err := indexWordlist(wordlist, word)
		if err != nil {
			return "", err
		}
		indexes = append(indexes, idx)
	}

	bits, err := bitsFromIndexes(indexes)
	if err != nil {
		return "", err
	}

	if len(bits) > 8448 {
		return "", errors.New("Invalid mnemonic")
	}

	dividerIdx := int(math.Floor(float64(len(bits))/33.0) * 32)
	entropy := bits[:dividerIdx]
	checksum := bits[dividerIdx:]

	entropyBuffer := []byte{}
	bins := chunkString(entropy, 8)
	for _, bin := range bins {
		i, err := strconv.ParseInt(bin, 2, 64)
		if err != nil {
			return "", err
		}
		entropyBuffer = append(entropyBuffer, byte(i))
	}
	newChecksum, err := checksumBits(entropyBuffer)
	if err != nil {
		return "", err
	}

	if newChecksum != checksum {
		return "", errors.New("Invalid mnemonic checksum")

	}

	return hex.EncodeToString(entropyBuffer), nil
}

func bitsFromIndexes(indexes []int) (string, error) {
	var buffer bytes.Buffer
	for _, idx := range indexes {
		_, err := buffer.WriteString(zeroPaddingBinaryFromInt(idx, 11))
		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}

func zeroPaddingBinaryFromInt(i int, length int) string {
	format := fmt.Sprintf("%%0%db", length)
	return fmt.Sprintf(format, i)
}

func indexWordlist(wordlist Wordlist, word string) (int, error) {
	for i, w := range wordlist {
		if w == word {
			return i, nil
		}
	}
	return -1, errors.New("Invalid mnemonic")
}

var defaultStrength = 128

// TODO: 名前変更
const eightBits = 8

var defaultRandomGenerator = RandomGeneratorImpl{}

func normalizeString(str string) string {
	return norm.NFKD.String(str)
}

func NewMnemonic(strength *int, generator RandomGenerator, wordlistLang string) (Mnemonic, error) {
	if strength == nil {
		strength = &defaultStrength
	}
	if generator == nil {
		generator = defaultRandomGenerator
	}

	var size int = *strength / eightBits
	hex := generator.Generate(size)
	e, err := NewEntropyFromBytes(hex)
	if err != nil {
		return Mnemonic{}, err
	}

	return e.ToMnemonic(wordlistLang)
}

type RandomGenerator interface {
	Generate(size int) []byte
}

type RandomGeneratorImpl struct {
}

func (rg RandomGeneratorImpl) Generate(size int) []byte {
	return []byte{}
}
