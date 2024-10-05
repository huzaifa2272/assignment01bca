package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
    Transaction   string
    Nonce         int
    PreviousHash  string
    Hash          string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
    block := &Block{transaction, nonce, previousHash, ""}
    block.Hash = CreateHash(block)
    return block
}

func CreateHash(block *Block) string {
    data := block.Transaction + block.PreviousHash + fmt.Sprintf("%d", block.Nonce)
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

func ListBlocks(chain []*Block) {
    for i, block := range chain {
        fmt.Printf("Block %d:\nTransaction: %s\nNonce: %d\nPrevious Hash: %s\nHash: %s\n\n", i, block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
    }
}

func ChangeBlock(block *Block, newTransaction string) {
    block.Transaction = newTransaction
    block.Hash = CreateHash(block)
}

func VerifyChain(chain []*Block) bool {
    for i := 1; i < len(chain); i++ {
        if chain[i].PreviousHash != chain[i-1].Hash {
            return false
        }
    }
    return true
}

func CalculateHash(stringToHash string) string {
    hash := sha256.Sum256([]byte(stringToHash))
    return hex.EncodeToString(hash[:])
}
