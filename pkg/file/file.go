package file

import (
	"os"
	"time"
)

type File struct {
	FilePath     string
	FileModified time.Time
	FileSize     int64
}

func New(path string, modified time.Time, size int64) File {
	return File{
		FilePath:     path,
		FileModified: modified,
		FileSize:     size,
	}
}

func (f *File) Path() string {
	return f.FilePath
}

func (f *File) Modified() time.Time {
	return f.FileModified
}

func (f *File) Size() int64 {
	return f.FileSize
}

func (f *File) Delete() error {
	return os.Remove(f.FilePath)
}
