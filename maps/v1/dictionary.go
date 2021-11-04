package maps

type Dictionary map[string]string

func (dict Dictionary) Search(word string) string {
	return dict[word]
}
