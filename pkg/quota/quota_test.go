package quota

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestQuota(t *testing.T) {
	err := os.Chdir("../../test")
	if err != nil {
		log.Fatal(err)
	}

	q := New(".", []string{".txt"}, 10)

	err = q.Walk()
	if err != nil {
		log.Fatal(err)
	}

	files := q.Candidates()
	assert.Equal(t, 1, len(files))
	assert.Equal(t, "folder_1/folder_2/folder_4/sauce.txt", files[0].FilePath)

	// Not gonna test delete- it should work right?!
}
