package main


type Dictionary map[string]string 

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

// We're implementing the `error` interface here.
// This helps distinguish this error from others.
// Relevant blog post: https://dave.cheney.net/2016/04/07/constant-errors
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// Search may return a value or an error, but we're really only
	// concerned with the error here. If there isn't an error, then
	// the word must exist in the dictionary.
	_, err := d.Search(word)

	switch err {
		case ErrNotFound:
			d[word] = definition
		case nil:
			return ErrWordExists
		default:
			return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist
		case nil:
			d[word] = definition
		default:
			return err
	}

	return nil
}