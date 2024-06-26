package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)


type Block struct{
	data 			map[string]interface{}
	hash 			string
	previousHash 	string
	timestamp 		int64
	pow 			int
}
type Blockchain struct{
	genesisBlock 	Block
	chain 			[]Block
	difficulty 		int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + strconv.FormatInt(b.timestamp, 10) + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)){
		b.pow++
		b.hash = b.calculateHash()
	}
}
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash: "0",
		timestamp: time.Now().Unix(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}
func (b *Blockchain) addBlock(from, to string, amount float64){
	blockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:  				blockData,
		previousHash: 		lastBlock.hash,
		timestamp: 			time.Now().Unix(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}
func (b Blockchain) isValid() bool {
	for i := 1 ; i < len(b.chain); i++  {
		previousBlock := b.chain[i-1]
		currentBlock := b.chain[i]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main () {
	// create new blockchain with mining difficulty of 2
	 blockchain := CreateBlockchain(2)

	// record transaction
	blockchain.addBlock("Alice", "Bob", 3)
	blockchain.addBlock("Max", "Bob", 5)
	blockchain.addBlock("Bob", "Alice", 12)
	blockchain.addBlock("Max", "Alice", 4)

	fmt.Println(blockchain.isValid())
	fmt.Println(blockchain)
}