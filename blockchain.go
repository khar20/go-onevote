package main

import (
	"fmt"
	"net/http"
	"onevote/models"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	Blockchain []models.Block
	mu         sync.Mutex
	voteQueue  []models.Vote
	batchSize  = 5
)

// process incoming votes
func ProcessVotes() {
	mu.Lock()
	defer mu.Unlock()

	if len(voteQueue) >= batchSize {

		newBlock := models.CreateBlock(voteQueue, &Blockchain)

		Blockchain = append(Blockchain, newBlock)

		voteQueue = []models.Vote{}
	}
}

// returns the current state of the blockchain
func GetChainHandler(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	return c.JSON(http.StatusOK, Blockchain)
}

// create a genesis block and inits the blockchain
func InitBlockChain() {
	mu.Lock()
	defer mu.Unlock()

	genesisBlock := models.Block{
		BlockNumber:  0,
		Timestamp:    time.Now(),
		PreviousHash: "",
		CurrentHash:  models.CalculateHash(fmt.Sprintf("%d%s", 0, time.Now().String())),
		Votes:        []models.Vote{},
	}
	Blockchain = append(Blockchain, genesisBlock)
}
