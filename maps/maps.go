package main

// ========================================================================
// make errors constant so they are immutable and reusable
// ========================================================================
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

func (d Dictionary) Search(word string) (string, error) {
	// maps now return 2 values. 2nd indicates whether it found a value sucecssfully
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist // not found in dictionary. cannot update
	case nil:
		d[word] = newDefinition // update dictionary value
	default:
		return err
	}
	return nil
}

// gross map caveats
// maps can be nil, which behaves like an empty map when reading. attempt to write to a nil map causes runtime panic
// never initialize a nil map variable
// e.g. var m map[string]string
// do this instead:
// var ditionary = map[string]string{} <-- the differnce is  {} and =
// OR
// var dictionary = make(map[string]string)
