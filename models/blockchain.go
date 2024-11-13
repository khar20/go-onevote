package models

import (
	"context"
	"encoding/json"
	"fmt"
	"onevote/database"
	"time"
)

// block struct
type Block struct {
	BlockNumber  int       `json:"block_number"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	CurrentHash  string    `json:"current_hash"`
	Votes        []Vote    `json:"votes"`
	MerkleRoot   string    `json:"merkle_root"`
}

// vote struct
type Vote struct {
	Location    string    `json:"location_id"`
	Type        string    `json:"type"`
	CandidateID string    `json:"candidate_id"`
	Timestamp   time.Time `json:"timestamp"`
}

// merkleNode struct
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  string
}

// builds a merkle tree using votes
func BuildMerkleTree(votes []Vote) *MerkleNode {
	if len(votes) == 0 {
		return nil
	}

	var nodes []MerkleNode
	for _, vote := range votes {
		data := fmt.Sprintf("%s:%s:%s:%s", vote.Location, vote.Type, vote.CandidateID, vote.Timestamp)
		hash := CalculateHash(data)
		nodes = append(nodes, MerkleNode{Hash: hash})
	}

	for len(nodes) > 1 {
		var newLevel []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right = left
			}
			newHash := CalculateHash(left.Hash + right.Hash)
			parent := MerkleNode{Left: &left, Right: &right, Hash: newHash}
			newLevel = append(newLevel, parent)
		}
		nodes = newLevel
	}

	return &nodes[0]
}

// creates a new block from votes
func CreateBlock(votes []Vote, blockchain *[]Block) Block {
	var merkleRoot string
	if len(votes) > 0 {
		merkleRoot = BuildMerkleTree(votes).Hash
	} else {
		merkleRoot = ""
	}

	block := Block{
		BlockNumber:  len(*blockchain),
		Timestamp:    time.Now(),
		PreviousHash: "",
		Votes:        votes,
		MerkleRoot:   merkleRoot,
	}

	if len(*blockchain) > 0 {
		block.PreviousHash = (*blockchain)[len(*blockchain)-1].CurrentHash
	}

	block.CurrentHash = CalculateHash(fmt.Sprintf("%d%s%s", block.BlockNumber, block.Timestamp, merkleRoot))
	return block
}

// retrieves all the blocks
func GetBlocks() ([]Block, error) {
	conn, err := database.ConnectVotesDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT block_data FROM blocks")
	if err != nil {
		return nil, fmt.Errorf("failed to query blocks: %w", err)
	}
	defer rows.Close()

	var blocks []Block
	for rows.Next() {
		var blockData []byte
		if err := rows.Scan(&blockData); err != nil {
			return nil, fmt.Errorf("failed to scan block_data: %w", err)
		}

		var block Block
		if err := json.Unmarshal(blockData, &block); err != nil {
			return nil, fmt.Errorf("failed to unmarshal block_data into Block: %w", err)
		}

		blocks = append(blocks, block)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return blocks, nil
}

// inserts a block into the database
func InsertBlock(block Block) error {
	blockData, err := json.Marshal(block)
	if err != nil {
		return fmt.Errorf("failed to marshal block: %w", err)
	}

	conn, err := database.ConnectVotesDB()
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `
		INSERT INTO blocks (block_data) 
		VALUES ($1)`, blockData)
	if err != nil {
		return fmt.Errorf("failed to insert block into database: %w", err)
	}

	return nil
}
