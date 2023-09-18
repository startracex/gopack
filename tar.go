package gopcak

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

type Tar struct {
	TarWriter *tar.Writer
}

// NewTar create new *Tar, need output target
func NewTar(target io.Writer, readers ...io.Reader) (p *Tar, e error) {
	p = &Tar{
		TarWriter: tar.NewWriter(target),
	}
	if p.TarWriter == nil {
		e = ErrNilTarget
		return
	}
	e = p.Add(readers...)
	return
}

// Add new readers, copy to TarWriter
func (t *Tar) Add(readers ...io.Reader) error {
	for _, src := range readers {
		if osFile, ok := src.(*os.File); ok {
			if StatResetName, err := osFile.Stat(); err != nil || StatResetName.IsDir() {
				err := filepath.Walk(osFile.Name(), func(path string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}
					file, err := os.Open(path)
					if err != nil {
						return err
					}
					return t.CopyFile(file)
				})
				if err != nil {
					return err
				}
			} else {
				_ = t.CopyFile(osFile)
			}
		}
		// TODO future
		// if _, ok := src.(_); ok {
		// }
	}
	return nil
}

func (t *Tar) Pack(readers ...io.Reader) error {
	err := t.Add(readers...)
	if err != nil {
		return err
	}
	return t.Close()
}

// Close to execute
func (t *Tar) Close() error {
	return t.TarWriter.Close()
}

// CopyFile copy a file to TarWriter
func (t *Tar) CopyFile(file *os.File) error {
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	h, err := tar.FileInfoHeader(fi, "")
	h.Name = file.Name()
	if err != nil {
		return err
	}
	err = t.TarWriter.WriteHeader(h)
	if err != nil {
		return err
	}
	return Copy(t.TarWriter, file)
}
