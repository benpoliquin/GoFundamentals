package main

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

// Search matches the word against the dictionary passed
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add word to the dictionary
func (d Dictionary) Add(word, definition string) error {
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

// Update the dictionary definition of a word
func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}

// Delete a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
