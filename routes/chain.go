package routes

import (
	"fmt"
	"net/http"
	"onevote/models"
	"time"

	"github.com/labstack/echo/v4"
)

// handlers

// todo
//func GetChainPage(c echo.Context) error {
//	blocks, err := models.GetBlocks()
//
//	if err != nil {
//		return err
//	}
//
//	data := templates.ChainData{
//		Blocks: blocks,
//	}
//
//	return Render(c, http.StatusOK, templates.ChainTempl(data))
//}

// returns the current state of the blockchain
func GetChain(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	return c.JSON(http.StatusOK, Blockchain)
}

// create a genesis block and inits the blockchain
func InitBlockChain() {
	genesisBlock := models.Block{
		BlockNumber:  0,
		Timestamp:    time.Now(),
		PreviousHash: "",
		CurrentHash:  models.CalculateHash(fmt.Sprintf("%d%s", 0, time.Now().String())),
		Votes:        []models.Vote{},
	}
	Blockchain = append(Blockchain, genesisBlock)
}
