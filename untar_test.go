package gopcak

import (
	"os"
	"testing"
)

func TestUnTar(t *testing.T) {
	PackTo, _ := os.Open("out")
	TarFile, _ := os.Open("tar.tar")
	z := NewUnZip()
	z.ReadFile(TarFile)
	z.Unpack(PackTo)
}
