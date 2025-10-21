// 代码生成时间: 2025-10-21 13:27:52
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Define the Block struct representing a block in the blockchain
type Block struct {
    gorm.Model
    Index     int    `gorm:"primaryKey"`
    Timestamp int64  "json:"timestamp"`
    Data      string `json:"data"`
    PrevHash  string `json:"prevHash"`
    Hash      string `json:"hash"`
}

// Blockchain represents the entire blockchain
type Blockchain struct {
    Blocks []*Block
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() *Blockchain {
    return &Blockchain{Blocks: make([]*Block, 0)}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) *Block {
    var lastBlock *Block
    if len(bc.Blocks) > 0 {
        lastBlock = bc.Blocks[len(bc.Blocks)-1]
    }

    newBlock := &Block{
        Index:     len(bc.Blocks) + 1,
        Timestamp: time.Now().Unix(),
        Data:      data,
        PrevHash: lastBlock.Hash,
    }

    newBlock.Hash = calculateHash(newBlock)
    bc.Blocks = append(bc.Blocks, newBlock)
    return newBlock
}

// calculateHash calculates the hash of a block
func calculateHash(block *Block) string {
    blockData := fmt.Sprintf("%d%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce)
    return fmt.Sprintf("%x", sha256.Sum256([]byte(blockData)))
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        currentBlock := bc.Blocks[i]
        previousBlock := bc.Blocks[i-1]
        if currentBlock.Hash != calculateHash(currentBlock) {
            return false
        }
        if currentBlock.PrevHash != previousBlock.Hash {
            return false
        }
    }
    return true
}

// ConnectToDatabase connects to the SQLite database
func ConnectToDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("blockchain.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return db
}

// SaveBlockchain saves the blockchain to the database
func SaveBlockchain(db *gorm.DB, bc *Blockchain) {
    db.AutoMigrate(&Block{})
    db.CreateInBatches(bc.Blocks, len(bc.Blocks))
}

// LoadBlockchain loads the blockchain from the database
func LoadBlockchain(db *gorm.DB) *Blockchain {
    var blocks []Block
    db.Find(&blocks)
    return &Blockchain{Blocks: blocks}
}

func main() {
    db := ConnectToDatabase()
    bc := LoadBlockchain(db)

    if !bc.IsValid() {
        fmt.Println("Blockchain is not valid")
        return
    }

    bc.AddBlock("Genesis Block")
    SaveBlockchain(db, bc)
    fmt.Println("Blockchain saved to database")
}
