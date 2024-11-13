package routes

import (
	"log"
	"net/http"
	"onevote/models"
	"onevote/templates"
	"strconv"
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

func PostVote(c echo.Context) error {
	candidateID := c.FormValue("candidate")
	if candidateID == "" {
		return c.String(http.StatusBadRequest, "Candidato no seleccionado")
	}

	candidate, err := models.GetCandidateByID(candidateID)
	if err != nil {
		return c.String(http.StatusBadRequest, "Candidato no seleccionado")
	}

	vote := models.Vote{
		Location:    "LL",
		Type:        "Consejo Nacional",
		CandidateID: strconv.Itoa(candidate.ID),
		Timestamp:   time.Now(),
	}

	mu.Lock()
	voteQueue = append(voteQueue, vote)
	mu.Unlock()

	if len(voteQueue) >= batchSize {
		go ProcessVotes()
	}

	return c.JSON(http.StatusAccepted, "Vote received and is being processed")
}

// creates a new block and insert it
func ProcessVotes() {
	mu.Lock()
	defer mu.Unlock()

	if len(voteQueue) >= batchSize {
		newBlock := models.CreateBlock(voteQueue, &Blockchain)

		Blockchain = append(Blockchain, newBlock)

		if err := models.InsertBlock(newBlock); err != nil {
			log.Printf("Failed to insert block into database: %v", err)
			return
		}

		voteQueue = []models.Vote{}
	}
}
