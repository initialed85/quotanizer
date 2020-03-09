package sorter

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"sort"
)

func Sort(files []file.File) []file.File {
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Modified().Unix() < files[j].Modified().Unix()
	})

	return files
}
