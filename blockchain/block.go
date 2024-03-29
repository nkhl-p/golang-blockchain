package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash: []byte{},
		Data: []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()

	return block
}

func (b *Block) DeriveHash()  {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func InitBlockChain() *BlockChain {
	return &BlockChain{Blocks:[]*Block{Genesis()}}
}