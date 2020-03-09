package file

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

const testFilePath = "some_file.txt"

func TestFile(t *testing.T) {
	err := os.Chdir("../../test")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(testFilePath, []byte("some data"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	f := New(testFilePath, time.Time{}, 640)

	assert.Equal(t, testFilePath, f.Path())

	err = f.Delete()
	if err != nil {
		log.Fatal(err)
	}

	err = f.Delete()
	assert.NotNil(t, err)
}
