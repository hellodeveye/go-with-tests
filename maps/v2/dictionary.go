package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
