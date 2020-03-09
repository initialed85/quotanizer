package walker

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestWalker(t *testing.T) {
	err := os.Chdir("../../test")
	if err != nil {
		log.Fatal(err)
	}

	w := New(".", []string{".txt"})

	files, err := w.Walk()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "folder_1/folder_2/folder_3/cheese.txt", files[0].Path())
	assert.Greater(t, files[0].Size(), int64(0))

	assert.Equal(t, "folder_1/folder_2/folder_4/sauce.txt", files[1].Path())
	assert.Greater(t, files[1].Size(), int64(0))
}
