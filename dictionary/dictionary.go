package dictionary

type Dictionary map[string]string

const (
	ErrNotFound       = DictionaryErr("could not find the word you were looking for")
	ErrWordExisted    = DictionaryErr("word is existed")
	ErrWordNotExisted = DictionaryErr("word is not existed")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if ok {
		return definition, nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return ErrWordExisted
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrWordNotExisted
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrWordNotExisted
	}
	delete(d, word)
	return nil
}
