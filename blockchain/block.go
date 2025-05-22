package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
}

func (b *Block) CalculateHash() []byte {
	var txBuf bytes.Buffer
	enc := gob.NewEncoder(&txBuf)
	_ = enc.Encode(b.Transactions)

	header := bytes.Join(
		[][]byte{
			IntToBytes(int64(b.Index)),
			IntToBytes(b.Timestamp),
			txBuf.Bytes(),
			b.PrevHash,
			IntToBytes(int64(b.Nonce)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(header)
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

func IntToBytes(num int64) []byte {
	buf := new(bytes.Buffer)
	_ = gob.NewEncoder(buf).Encode(num)
	return buf.Bytes()
}
