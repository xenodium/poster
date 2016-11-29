package source

import (
	"fmt"
	"io/ioutil"
	"nfl/strings"
	"os"
	"path/filepath"
)

type Iterator struct {
	filePaths []string
	pos       int

	textIter *strings.Iterator
}

func NewIterator(filePaths []string) *Iterator {
	it := Iterator{
		filePaths: filePaths,
		pos:       -1,
	}
	return &it
}

func (it *Iterator) Count() int {
	return len(it.filePaths)
}

func (it *Iterator) Reset() {
	it.pos = -1
	it.textIter = nil
}

func (it *Iterator) Done() bool {
	return it.pos+1 >= it.Count() && it.textIter != nil && it.textIter.Done()
}

func (it *Iterator) Next() {
	if it.Done() {
		return
	}
	it.pos++
	it.textIter = iterForFilePath(it.filePaths[it.pos])
}

// func (it *Iterator) FilteredRead() *string {
// 	if it.Done() {
// 		return nil
// 	}
// 	if it.textIter == nil {
// 		it.Next()
// 	}
// 	if it.textIter.Done() {
// 		it.Next()
// 	}
// 	return it.textIter.Read()
// }

func (it *Iterator) Read() *string {
	if it.Done() {
		return nil
	}
	if it.textIter == nil {
		it.Next()
	}
	if it.textIter.Done() {
		it.Next()
	}
	return it.textIter.Read()
}

func iterForFilePath(filePath string) *strings.Iterator {
	textBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("could not read file: %v, %v", filePath, err))
	}
	return strings.NewIterator(string(textBytes))
}

func ListFiles(rootPath string) []string {
	var filePaths = []string{}
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		filePaths = append(filePaths, path)
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("Can't list sources: %v", err))
	}

	return filePaths
}

// func (it *Iterator) Done() bool {
// 	if it.pos >= len(it.filePaths) {
// 		return true
// 	}
// 	if it.pos == -1 {
// 		return false
// 	}
// 	return it.fileReader.Done()
// }

// func (it *Iterator) nextFileReader() {
// 	it.pos++
// 	// if it.Done() {
// 	// 	return
// 	// }
// 	fr, err := newFileReader(it.filePaths[it.pos])
// 	if err != nil {
// 		return
// 	}
// 	it.fileReader = fr
// }

// func (it *Iterator) Read() *string {
// 	if it.Done() {
// 		return nil
// 	}
// 	if it.fileReader.Done() {
// 		it.nextFileReader()
// 	}
// 	return it.fileReader.Read()
// }
