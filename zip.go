package gopcak

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
)

type Zip struct {
	ZipWriter *zip.Writer
}

// NewZip Create new *Zip, need output target
func NewZip(target io.Writer, readers ...io.Reader) (p *Zip, e error) {
	p = &Zip{
		ZipWriter: zip.NewWriter(target),
	}
	if p.ZipWriter == nil {
		e = ErrNilTarget
		return
	}
	e = p.Add(readers...)
	return
}

// Add new readers, copy to ZipWriter
func (z *Zip) Add(readers ...io.Reader) error {
	for _, src := range readers {
		if osFile, ok := src.(*os.File); ok {
			fi, _ := osFile.Stat()
			if !fi.IsDir() {
				err := z.CopyFile(osFile)
				if err != nil {
					return err
				}
			} else {
				err := filepath.Walk(osFile.Name(), func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() {
						return nil
					}
					file, _ := os.Open(path)
					return z.CopyFile(file)
				})
				if err != nil {
					return err
				}
			}
		}
		// TODO future
		// if _, ok := src.(_); ok {
		// }
	}
	return nil
}

func (z *Zip) Pack(readers ...io.Reader) error {
	err := z.Add(readers...)
	if err != nil {
		return err
	}
	return z.Close()
}

// Close to execute
func (z *Zip) Close() error {
	return z.ZipWriter.Close()
}

// CopyFile copy a file to ZipWriter
func (z *Zip) CopyFile(file *os.File) error {
	zipWriter := z.ZipWriter
	name := path.Clean(file.Name())
	made, err := zipWriter.Create(name)
	if err != nil {
		return err
	}
	return Copy(made, file)
}
