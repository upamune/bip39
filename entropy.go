package bip39

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/pkg/errors"
	"strings"
	"strconv"
)

type Entropy struct {
	// TODO: 名前変更 → buffer ?
	Hex string
}

func (e Entropy) ToMnemonic(wordlistLang string) (Mnemonic, error) {
	entropyBuffer := []byte(e.Hex)
	l := len(entropyBuffer)
	if l == 0 || l > 1024 || l%4 != 0 {
		return Mnemonic{}, errors.New("Invalid entropy")
	}

	entropyBits, err := bytesToBinary(entropyBuffer)
	if err != nil {
		return Mnemonic{}, err
	}

	checksum, err := checksumBits(entropyBuffer)
	if err != nil {
		return Mnemonic{}, err
	}

	bits := entropyBits + checksum
	chunks := chunkString(bits, 10)
	words, err := chunkToWords(chunks, wordlistLang)
	if err != nil {
		return Mnemonic{}, err
	}

	joint := " "
	if wordlistLang == "japanese" {
		joint = "　"
	}

	return Mnemonic{Words: strings.Join(words, joint)}, nil
}

func chunkToWords(chunks []string, wordlistLang string) ([]string, error) {
	wordlist, ok := GetWordlists(wordlistLang)
	if !ok {
		return []string{}, errors.New("no such a language wordlist")
	}
	words := []string{}
	for _, w := range chunks {
		idx, err := strconv.Atoi(w)
		if err != nil {
			return []string{}, err
		}

		if idx >= len(wordlist) {
			return []string{}, errors.New("out of range wordlist")
		}
		words = append(words, wordlist[idx])
	}

	return words, nil
}

func chunkString(str string, length int) []string {
	a := []rune(str)
	s := []string{}
	res := ""
	lastIdx := 0
	for i, r := range a {
		res = res + string(r)
		if i > 0 && (i+1)%length == 0 {
			s = append(s, res)
			res = ""
			lastIdx = i
		}
	}
	l := len(a) - 1
	if l != lastIdx {
		s = append(s, string(a[lastIdx+1:]))
	}

	return s
}

func checksumBits(buf []byte) (string, error) {
	hash := sha256.New().Sum(buf)

	var ENT = len(buf) * eightBits
	var CS = ENT / 32

	return bytesToBinary(hash)[:CS]
}

func bytesToBinary(b []byte) (string, error) {
	var buffer bytes.Buffer

	for _, b := range b {
		_, err := buffer.WriteString(zeroPadding(b, eightBits))
		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}

func zeroPadding(b byte, length int) string {
	format := fmt.Sprintf("%%0%db", length)
	return fmt.Sprintf(format, b)
}
