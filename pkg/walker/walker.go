package walker

import (
	"github.com/initialed85/quotanizer/pkg/file"
	"os"
	"path/filepath"
	"strings"
)

type Walker struct {
	path     string
	suffixes []string
	files    []file.File
}

func New(path string, suffixes []string) Walker {
	lowerSuffixes := make([]string, 0)
	for _, suffix := range suffixes {
		lowerSuffixes = append(
			lowerSuffixes,
			"."+strings.ToLower(strings.TrimLeft(suffix, ".")),
		)
	}

	return Walker{
		path:     path,
		suffixes: suffixes,
	}
}

func (w *Walker) walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	ext := strings.ToLower(filepath.Ext(path))

	if len(w.suffixes) > 0 {
		ignore := true
		for _, s := range w.suffixes {
			if s == ext {
				ignore = false

				break
			}
		}

		if ignore {
			return nil
		}
	}

	if !info.Mode().IsRegular() {
		return nil
	}

	w.files = append(
		w.files,
		file.New(
			path,
			info.ModTime(),
			info.Size(),
		),
	)

	return nil
}

func (w *Walker) Walk() ([]file.File, error) {
	w.files = make([]file.File, 0)
	err := filepath.Walk(w.path, w.walkFunc)
	if err != nil {
		return []file.File{}, err
	}

	return w.files, err
}
