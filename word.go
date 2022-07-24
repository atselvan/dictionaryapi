package dictionaryapi

import (
	"fmt"

	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	wordEntriesApiPath = "/api/v2/entries/en/%s"
)

type (
	// Word represents a Word who's meaning is fetched from the dictionary API.
	Word struct {
		Word      string      `json:"word"`
		Phonetic  string      `json:"phonetic"`
		Phonetics []Phonetics `json:"phonetics"`
		Origin    string      `json:"origin"`
		Meanings  []Meanings  `json:"meanings"`
	}

	// Phonetics represents the sound of the Word.
	Phonetics struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	}

	// Definitions represent a statement of the exact meaning of the Word.
	Definitions struct {
		Definition string        `json:"definition"`
		Example    string        `json:"example"`
		Synonyms   []interface{} `json:"synonyms"`
		Antonyms   []interface{} `json:"antonyms"`
	}

	// Meanings represents the different meanings of the Word
	Meanings struct {
		PartOfSpeech string        `json:"partOfSpeech"`
		Definitions  []Definitions `json:"definitions"`
	}

	WordsManager interface {
		Get(string) (*Word, *errors.RestErr)
	}

	wordManager struct {
		Client *Client
	}
)

// Get queries the dictionary API to fetch the meaning of a word which is passed as input.
// The method returns the meaning of the word if the request is successful.
// The method returns a error if
//	- the API request fails.
//	- the word is not found.
func (wm *wordManager) Get(word string) (*Word, *errors.RestErr) {
	result := new([]Word)
	err := wm.Client.get(fmt.Sprintf(wordEntriesApiPath, word), result)
	if err != nil {
		return nil, err
	}
	return &(*result)[0], nil
}
