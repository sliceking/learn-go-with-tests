package main

const (
	ErrMissingWord = DictionaryErr("could not find the word you were looking for")
	ErrWordExists  = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")

)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	w, ok := d[word]
	if !ok {
		return "", ErrMissingWord
	}
	return w, nil
}

// Add adds a word to the dictionary.
func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrMissingWord:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update updates a word in the dictionary.
func (d Dictionary) Update(word string, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrMissingWord:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete updates a word in the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
