package maps

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dict Dictionary) Add(word, definition string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Update(word, definition string) error {
	_, err := dict.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[word] = definition
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Delete(word string) {
	delete(dict, word)
}
