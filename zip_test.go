package gopcak

import (
	"os"
	"testing"
)

func TestZip(*testing.T) {
	ZipFile, _ := os.Create("zip.zip")
	PackFrom, _ := os.Open("test")
	z, _ := NewZip(ZipFile)
	z.Pack(PackFrom)
}
