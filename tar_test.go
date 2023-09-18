package gopcak

import (
	"os"
	"testing"
)

func TestTar(*testing.T) {
	TarFile, _ := os.Create("tar.tar")
	PackFrom, _ := os.Open("test")
	z, _ := NewTar(TarFile)
	z.Pack(PackFrom)
}
