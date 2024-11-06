package models

import (
	"fmt"
	"time"
)

// block struct
type Block struct {
	BlockNumber  int       `json:"block_number"`
	PreviousHash string    `json:"previous_hash"`
	Timestamp    time.Time `json:"timestamp"`
	Votes        []Vote    `json:"votes"`
	CurrentHash  string    `json:"current_hash"`
}

// vote struct
type Vote struct {
	UserID      string    `json:"voter_id"`
	SedeID      string    `json:"sede_id"`
	Type        string    `json:"tipo_voto"`
	CandidateID string    `json:"candidato_id"`
	Timestamp   time.Time `json:"timestamp"`
}

// merklenode struct
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
		data := fmt.Sprintf("%s:%s:%s:%s:%s", vote.UserID, vote.SedeID, vote.Type, vote.CandidateID, vote.Timestamp)
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
	}

	if len(*blockchain) > 0 {
		block.PreviousHash = (*blockchain)[len(*blockchain)-1].CurrentHash
	}
	block.CurrentHash = CalculateHash(fmt.Sprintf("%d%s%s", block.BlockNumber, block.Timestamp, merkleRoot))
	return block
}

// retrieves all blocks from the database
//func GetBlocks(db *sql.DB) ([]Block, error) {
//	query := `SELECT block_number, previous_hash, timestamp, votes, current_hash FROM blocks ORDER BY block_number ASC`
//
//	rows, err := db.Query(query)
//	if err != nil {
//		return nil, fmt.Errorf("error querying blocks: %v", err)
//	}
//	defer rows.Close()
//
//	var blocks []Block
//
//	for rows.Next() {
//		var block Block
//		var votesJSON string
//
//		if err := rows.Scan(&block.BlockNumber, &block.PreviousHash, &block.Timestamp, &votesJSON, &block.CurrentHash); err != nil {
//			return nil, fmt.Errorf("error scanning row: %v", err)
//		}
//
//		if err := json.Unmarshal([]byte(votesJSON), &block.Votes); err != nil {
//			return nil, fmt.Errorf("error unmarshaling votes: %v", err)
//		}
//
//		blocks = append(blocks, block)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("error during rows iteration: %v", err)
//	}
//
//	return blocks, nil
//}
