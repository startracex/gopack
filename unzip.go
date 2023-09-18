package gopcak

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
)

type UnZip struct {
	ZipReader *zip.Reader
}

func NewUnZip() *UnZip {
	return &UnZip{}
}

//Read create reader from io.ReaderAt, size
func (z *UnZip) Read(r io.ReaderAt, size int64) error {
	reader, err := zip.NewReader(r, size)
	if err != nil {
		return err
	}
	z.ZipReader = reader
	return nil
}

// ReadFile read a zip file
func (z *UnZip) ReadFile(file *os.File) error {
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	reader, err := zip.NewReader(file, fi.Size())
	z.ZipReader = reader
	return nil
}

// Unpack files to target
func (z *UnZip) Unpack(target io.Writer) error {
	for _, file := range z.ZipReader.File {
		if osFile, ok := target.(*os.File); ok {
			fi := file.FileInfo()
			name := path.Join(osFile.Name(), fi.Name())
			name = filepath.ToSlash(name)
			out, err := CreateFile(name, fi.Mode())
			if err != nil {
				return err
			}
			src, err := file.Open()
			if err != nil {
				return err
			}
			err = Copy(out, src)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
