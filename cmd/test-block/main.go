package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/wisepythagoras/dimoschain/dimos"
)

func main() {
	// Load the database.
	blockchain, err := dimos.InitChainDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// First we need to get the current block. This will determine which IDx this
	// new block will take, as well as the previous hash.
	currentBlock, err := blockchain.GetCurrentBlock()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a test transaction.
	tx := dimos.Transaction{
		Hash:      nil,
		Amount:    0.001,
		From:      []byte("test1"),
		To:        []byte("test2"),
		Signature: []byte("test1signature"),
	}

	// Construct the new block.
	block := dimos.Block{
		IDx:          currentBlock.IDx + 1,
		Hash:         nil,
		PrevHash:     currentBlock.Hash,
		MerkleRoot:   nil,
		Timestamp:    time.Now().Unix(),
		Transactions: []dimos.Transaction{},
		Signature:    []byte("test"),
	}

	// Adding a transaction will also change the hash on the block.
	block.AddTransaction(&tx)

	// But we run this anyway, since a block could be empty.
	block.UpdateHash()

	// Finally add the block.
	success, err := blockchain.AddBlock(&block)

	if err != nil {
		log.Fatal(err)
	} else if !success {
		log.Fatal("Unale to add the new block")
	} else {
		hexHash := hex.EncodeToString(block.Hash)
		msg := fmt.Sprintf("The block was added with hash %s", hexHash)
		log.Println(msg)
	}
}