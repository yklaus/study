package dictionary

import (
	"errors"
	"fmt"
)

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-exists word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}

	return "", errNotFound
}

// Add word in Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) String() string {
	str := ""

	for word, def := range d {
		str += fmt.Sprint("word:", word, ", definition: ", def, "\n")
	}

	return str
}
