package routes

import (
	"net/http"
	"onevote/models"
	"onevote/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

// handlers
func GetCandidatesPage(c echo.Context) error {
	candidates, err := models.GetCandidates()

	if err != nil {
		return err
	}

	data := templates.CandidatesData{
		Candidates: candidates,
	}

	return Render(c, http.StatusOK, templates.CandidatesTempl(data))
}

func GetCandidateProfile(c echo.Context) error {
	candidateId, err := strconv.Atoi(c.Param("candidate-id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	data := templates.CandidateProfileData{
		Candidate: nil,
	}

	candidates, err := models.GetCandidates()
	if err != nil {
		return err
	}

	for _, candidate := range candidates {
		if candidate.ID == candidateId {
			data.Candidate = &candidate
			break
		}
	}

	if data.Candidate == nil {
		return c.NoContent(http.StatusOK)
	}

	return Render(c, http.StatusOK, templates.CandidateProfileTempl(data))
}
