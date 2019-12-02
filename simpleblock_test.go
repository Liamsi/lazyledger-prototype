package lazyledger

import (
	"testing"
)

func TestSimpleBlock(t *testing.T) {
	sb := NewSimpleBlock([]byte{0})

	sb.AddMessage(*NewMessage([namespaceSize]byte{0}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{1}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{1}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{1}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{3}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{3}, []byte("foo")))
	sb.AddMessage(*NewMessage([namespaceSize]byte{4}, []byte("foo")))

	proofStart, proofEnd, proof, messages, hashes := sb.(*SimpleBlock).ApplicationProof([namespaceSize]byte{1})
	if messages == nil {
		t.Error("ApplicationProof incorrectly returned no messages")
	}
	result := sb.(*SimpleBlock).VerifyApplicationProof([namespaceSize]byte{1}, proofStart, proofEnd, proof, messages, hashes)
	if !result {
		t.Error("VerifyApplicationProof incorrectly returned false")
	}

	proofStart, proofEnd, proof, messages, hashes = sb.(*SimpleBlock).ApplicationProof([namespaceSize]byte{1})
	proof[0][0] = 0xFF
	if messages == nil {
		t.Error("ApplicationProof incorrectly returned no messages")
	}
	result = sb.(*SimpleBlock).VerifyApplicationProof([namespaceSize]byte{1}, proofStart, proofEnd, proof, messages, hashes)
	if result {
		t.Error("VerifyApplicationProof incorrectly returned true")
	}

	proofStart, proofEnd, proof, messages, hashes = sb.(*SimpleBlock).ApplicationProof([namespaceSize]byte{2})
	if messages != nil {
		t.Error("ApplicationProof incorrectly returned messages")
	}
	result = sb.(*SimpleBlock).VerifyApplicationProof([namespaceSize]byte{2}, proofStart, proofEnd, proof, messages, hashes)
	if !result {
		t.Error("VerifyApplicationProof incorrectly returned false")
	}

	proofStart, proofEnd, proof, messages, hashes = sb.(*SimpleBlock).ApplicationProof([namespaceSize]byte{2})
	proof[0][0] = 0xFF
	if messages != nil {
		t.Error("ApplicationProof incorrectly returned messages")
	}
	result = sb.(*SimpleBlock).VerifyApplicationProof([namespaceSize]byte{2}, proofStart, proofEnd, proof, messages, hashes)
	if result {
		t.Error("VerifyApplicationProof incorrectly returned true")
	}
}
