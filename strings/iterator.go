package strings

import "regexp"

type Iterator struct {
	content string
	pos     int
}

func NewIterator(content string) *Iterator {
	return &Iterator{
		// Remove all whitespace. Makes for better visuals.
		content: regexp.MustCompile(`\s+`).ReplaceAllString(content, ` `),
		pos:     -1,
	}
}

func (it *Iterator) Count() int {
	return len(it.content)
}

func (it *Iterator) Done() bool {
	return it.pos+1 >= it.Count()
}

func (it *Iterator) Next() {
	if it.Done() {
		return
	}
	it.pos++
}

func (it *Iterator) Read() *string {
	if it.Done() {
		return nil
	}
	it.Next()
	str := string(it.content[it.pos])
	return &str
}
