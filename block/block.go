package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/spiermar/zerocoin/proto"
)

func calculateHash(index uint64, previousHash string, timestamp int64, data string) string {
	var b bytes.Buffer
	b.WriteString(strconv.FormatUint(index, 10))
	b.WriteString(previousHash)
	b.WriteString(strconv.FormatInt(timestamp, 10))
	b.WriteString(data)
	h := sha256.New()
	h.Write([]byte(b.String()))
	s := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return s
}

func calculateHashForBlock(block *proto.Block) string {
	return calculateHash(block.Index, block.PreviousHash, block.Timestamp, block.Data)
}

// NewBlock returns a new Block
func NewBlock(index uint64, previousHash string, data string) *proto.Block {
	n := time.Now()
	b := new(proto.Block)
	b.Index = index
	b.Hash = calculateHash(index, previousHash, n.Unix(), data)
	b.PreviousHash = previousHash
	b.Timestamp = n.UnixNano()
	b.Data = data
	return b
}

// IsValidBlock validates a block against a previous block
func IsValidBlock(block *proto.Block, previousBlock *proto.Block) bool {
	if previousBlock.Index+1 != block.Index {
		return false
	} else if previousBlock.Hash != block.PreviousHash {
		return false
	} else if calculateHashForBlock(block) != block.Hash {
		return false
	}
	return true
}
