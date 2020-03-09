package sizer

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"github.com/initialed85/quotanizer/pkg/sorter"
)

func GetTotal(files []file.File) int64 {
	var total int64 = 0
	for _, f := range files {
		total += f.Size()
	}

	return total
}

func GetCandidates(files []file.File, quota int64) []file.File {
	total := GetTotal(files)

	files = sorter.Sort(files)

	var progress int64 = 0
	candidates := make([]file.File, 0)

	for _, f := range files {
		if total-progress <= quota {
			return candidates
		}

		progress += f.Size()
		candidates = append(candidates, f)
	}

	// if we got here we deleted everything; oh well, you asked for this!
	return candidates
}
