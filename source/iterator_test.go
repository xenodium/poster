package source

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestNew(t *testing.T) {
	iter := New([]string{"path/one", "path/two"})
	if got, want := len(iter.filePaths), 2; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRead(t *testing.T) {
	testDir := newTestDir()
	iter := New([]string{
		writeTextToFilePath("hello", path.Join(testDir, "hello.txt")),
		writeTextToFilePath("world", path.Join(testDir, "world.txt")),
	})

	// tests := []struct {
	// 	want string
	// }{
	// 	{
	// 		want: "h",
	// 	},
	// }
	// for test := range tests {
	// }
	if got, want := iter.Read(), "h"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "e"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "l"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "l"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "o"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}

	if got, want := iter.Read(), "w"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "o"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "r"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "l"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}
	if got, want := iter.Read(), "d"; *got != want {
		t.Errorf("got %v want %v", *got, want)
	}

	if got, want := iter.Done(), true; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func newTestDir() string {
	testDir, err := ioutil.TempDir("", "TestRead")
	if err != nil {
		panic(err)
	}
	return testDir
}
func writeTextToFilePath(text string, filePath string) string {
	err := ioutil.WriteFile(filePath, []byte(text), 0755)
	if err != nil {
		panic(err)
	}
	return filePath
}
