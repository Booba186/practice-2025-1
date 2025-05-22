package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func (b *Block) CalculateHash() []byte {
	headers := bytes.Join(
		[][]byte{
			[]byte(string(b.Index)),
			[]byte(string(b.Timestamp)),
			[]byte(b.Data),
			b.PrevHash,
			[]byte(string(b.Nonce)),
		}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

func (b *Block) Mine(difficulty int) {
	prefix := bytes.Repeat([]byte{0}, difficulty)
	for {
		b.Hash = b.CalculateHash()
		if bytes.HasPrefix(b.Hash, prefix) {
			break
		}
		b.Nonce++
	}
}
