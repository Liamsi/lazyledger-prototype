package lazyledger

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestFlagHasher(t *testing.T) {
	ndf := NewNamespaceDummyFlagger()
	fh := NewFlagHasher(ndf, sha256.New())
	data := make([]byte, 100)
	rand.Read(data)

	fh.Write([]byte{0})
	fh.Write(data)
	leaf1 := fh.Sum(nil)
	fh.Reset()
	if !bytes.Equal(ndf.NodeFlag(leaf1), ndf.LeafFlag(data)) {
		t.Error("flag for leaf node incorrect")
	}

	fh.Write([]byte{0})
	fh.Write(data)
	leaf2 := fh.Sum(nil)
	fh.Reset()
	if !bytes.Equal(ndf.NodeFlag(leaf2), ndf.LeafFlag(data)) {
		t.Error("flag for leaf node incorrect")
	}

	fh.Write([]byte{1})
	fh.Write(leaf1)
	fh.Write(leaf2)
	parent := fh.Sum(nil)
	fh.Reset()
	if !bytes.Equal(ndf.NodeFlag(parent), ndf.Union(leaf1, leaf2)) {
		t.Error("flag for parent node incorrect")
	}
}
