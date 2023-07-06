package docker

import (
	"bytes"
	"testing"
)

func TestInit(t *testing.T) {
	byte := []byte(`hello world


`)
	byte = bytes.TrimRight(byte, "\n")
	t.Log(byte)
}
