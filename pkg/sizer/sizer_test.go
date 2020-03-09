package sizer

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestFiles() []file.File {
	time1 := time.Time{}
	time2 := time1.Add(time.Second * -1)
	time3 := time2.Add(time.Second * -1)

	return []file.File{
		{"/path/file_2.txt", time2, 100},
		{"/path/file_1.txt", time1, 100},
		{"/path/file_3.txt", time3, 100},
	}
}

func TestGetTotal(t *testing.T) {
	assert.Equal(
		t,
		int64(300),
		GetTotal(getTestFiles()),
	)
}

func TestGetCandidates(t *testing.T) {
	files := getTestFiles()

	assert.Equal(
		t,
		[]file.File{
			{"/path/file_3.txt", files[2].Modified(), 100},
			{"/path/file_2.txt", files[0].Modified(), 100},
		},
		GetCandidates(files, 100),
	)
}
