package routes

import (
	"net/http"
	"onevote/models"
	"onevote/templates"
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

// handlers
func GetVotePage(c echo.Context) error {
	candidates, err := models.GetCandidates()

	if err != nil {
		return err
	}

	data := templates.VotingData{
		Candidates: candidates,
	}

	return Render(c, http.StatusOK, templates.VotingTempl(data))
}

// todo
func PostVote(c echo.Context) error { // htmx
	var msg models.Vote
	if err := c.Bind(&msg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid vote format"})
	}

	msg.Timestamp = time.Now()

	mu.Lock()
	voteQueue = append(voteQueue, msg)
	mu.Unlock()

	if len(voteQueue) >= batchSize {
		go ProcessVotes()
	}

	return c.JSON(http.StatusAccepted, map[string]string{"message": "Vote received and is being processed"})
}

// todo
func ProcessVotes() {
	mu.Lock()
	defer mu.Unlock()

	if len(voteQueue) >= batchSize {

		newBlock := models.CreateBlock(voteQueue, &Blockchain)

		Blockchain = append(Blockchain, newBlock)

		voteQueue = []models.Vote{}
	}
}
