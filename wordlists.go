package bip39

import (
	"encoding/json"
	"fmt"

	"github.com/rakyll/statik/fs"
	_ "github.com/upamune/bip39/statik"
)

type Wordlist []string

const DefaultWordlist = "english"
var languages = []string{
	"english",
	"french",
	"italian",
	"japanese",
	"spanish",
}
func GetLanguages() []string {
	return languages
}

var languageMap = make(map[string]struct{}, len(languages))
func IsValidLanguage(lang string) bool {
	_, ok := languageMap[lang]
	return ok
}

var wordDic = make(map[string]Wordlist, len(languages))
func GetWordlists(lang string) (Wordlist, bool) {
	ws, ok := wordDic[lang]
	return ws, ok
}


//go:generate statik -src=./wordlists
func init() {
	for _, l := range languages {
		languageMap[l] = struct{}{}
	}

	fs, err := fs.New()
	if err != nil {
		panic(err)
	}

	for _, language := range languages {
		filename := fmt.Sprintf("/%s.json", language)
		file, err := fs.Open(filename)
		if err != nil {
			panic(err)
		}
		var words Wordlist
		if err := json.NewDecoder(file).Decode(&words); err != nil {
			panic(err)
		}
		wordDic[language] = words
	}
}
