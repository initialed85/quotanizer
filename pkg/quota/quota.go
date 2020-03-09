package quota

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"github.com/initialed85/quotanizer/pkg/sizer"
	"github.com/initialed85/quotanizer/pkg/walker"
	"log"
)

type Quota struct {
	path        string
	quotaWalker walker.Walker
	quota       int64
	files       []file.File
}

func New(path string, suffixes []string, quota int64) Quota {
	log.Printf("new quota for path='%v', suffixes=%v, quota=%v GB", path, suffixes, quota/1e+9)

	return Quota{
		path:        path,
		quotaWalker: walker.New(path, suffixes),
		quota:       quota,
		files:       make([]file.File, 0),
	}
}

func (q *Quota) Walk() error {
	log.Printf("walking '%v'", q.path)

	files, err := q.quotaWalker.Walk()
	if err != nil {
		return err
	}

	q.files = files

	log.Printf(
		"got %v files at '%v' totalling %v GB",
		len(q.files),
		q.path,
		float64(sizer.GetTotal(q.files))/1e+9,
	)

	return nil
}

func (q *Quota) Candidates() []file.File {
	files := sizer.GetCandidates(q.files, q.quota)

	log.Printf(
		"got %v candidate files at '%v' totalling %v GB",
		len(files),
		q.path,
		float64(sizer.GetTotal(files))/1e+9,
	)

	return files
}

func (q *Quota) Delete(files []file.File) error {
	log.Printf(
		"deleting %v files at '%v' totalling %v GB",
		len(files),
		q.path,
		float64(sizer.GetTotal(files))/1e+9,
	)

	for _, f := range files {
		log.Printf("deleting '%v'", f.Path())

		err := f.Delete()
		if err != nil {
			return err
		}
	}

	return nil
}
