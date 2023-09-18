package gopcak

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
)

var (
	ErrNilTarget = errors.New("target is nil")
)

func Copy(target io.Writer, source io.Reader) error {
	_, err := io.Copy(target, source)
	return err
}

// CreateFile create files directory and file
func CreateFile(name string, mode fs.FileMode) (*os.File, error) {
	dir := path.Dir(name)
	err := os.MkdirAll(dir, mode)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
