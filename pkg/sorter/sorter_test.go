package sorter

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	time1 := time.Time{}
	time2 := time1.Add(time.Second * -1)
	time3 := time2.Add(time.Second * -1)

	files := []file.File{
		{"/path/file_1.txt", time1, 1},
		{"/path/file_2.txt", time2, 1},
		{"/path/file_3.txt", time3, 1},
	}

	assert.Equal(
		t,
		[]file.File{
			{"/path/file_3.txt", time3, 1},
			{"/path/file_2.txt", time2, 1},
			{"/path/file_1.txt", time1, 1},
		},
		Sort(files),
	)
}
