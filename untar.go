package gopcak

import (
	"archive/tar"
	"io"
	"os"
	"path"
	"path/filepath"
)

type UnTar struct {
	TarReader *tar.Reader
}

func NewUnTar() *UnTar {
	return &UnTar{}
}

// Read Set reader to TarReader
func (t *UnTar) Read(reader io.Reader) *UnTar {
	t.TarReader = tar.NewReader(reader)
	return t
}

// ReadFile is alias of Read
func (t *UnTar) ReadFile(file *os.File) *UnTar {
	return t.Read(file)
}

// Unpack files to target
func (t *UnTar) Unpack(target io.Writer) error {

	for hdr, err := t.TarReader.Next(); err != io.EOF; hdr, err = t.TarReader.Next() {
		if osFile, ok := target.(*os.File); ok {
			fi := hdr.FileInfo()
			name := path.Join(osFile.Name(), fi.Name())
			name = filepath.ToSlash(name)
			out, err := CreateFile(name, fi.Mode())
			if err != nil {
				return err
			}
			err = Copy(out, t.TarReader)
		}
	}
	return nil
}
