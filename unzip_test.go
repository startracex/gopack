package gopcak

import (
	"os"
	"testing"
)

func TestUnZip(*testing.T) {
	PackTo, _ := os.Open("out")
	ZipFile, _ := os.Open("zip.zip")
	z := NewUnZip()
	z.ReadFile(ZipFile)
	z.Unpack(PackTo)
}
