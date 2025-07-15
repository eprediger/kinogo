package infrastructure

import (
	application "application/use_cases/sources"
	domain "domain/model"
	"encoding/json"
	feedDto "infrastructure/http/feed/dto"
	problem "infrastructure/http/problem"
	"io"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed"
)

type SourceHandler interface {
	GetAllSources(w http.ResponseWriter, r *http.Request)
	CreateSource(w http.ResponseWriter, r *http.Request)
}

func NewSourcesHandler(service application.SourceService) SourceHandler {
	return &sourceHandler{
		service: service,
	}
}

type sourceHandler struct {
	service application.SourceService
}

// GetAllSources returns all existing sources.
func (h *sourceHandler) GetAllSources(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sources := h.service.GetAllSources(ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sources)
}

// CreateSource Save a new source from a RSS Feed URL.
func (h *sourceHandler) CreateSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var feedDto feedDto.FeedDto
	if err := json.Unmarshal(body, &feedDto); err != nil {
		problem.Write(
			w,
			problem.ProblemResponse{
				Type:   "Internal Server Error",
				Title:  "",
				Status: http.StatusBadRequest,
				Detail: "The request body contains invalid JSON",
			},
		)
		return
	}

	feed, err := domain.NewFeed(feedDto)
	if err != nil {
		errResponse := problem.ProblemResponse{
			Type:   "Request Error",
			Title:  "",
			Status: http.StatusUnprocessableEntity,
			Detail: strings.ReplaceAll(err.Error(), "\"", "'"),
		}
		problem.Write(w, errResponse)
		return
	}

	feedParser := gofeed.NewParser()

	parsedFeed, err := feedParser.ParseURL(feed.GetUrl())

	if err != nil {
		errResponse := problem.ProblemResponse{
			Type:   "Parse Error",
			Title:  "",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}
		problem.Write(w, errResponse)
		return
	}
	/* TODO: validation
	 * so far the only unique id seems to be the URL, also in feed.FeedLink
	 *	response feed already exists
	 */

	newSource := domain.NewSource(parsedFeed)
	h.service.CreateSource(ctx, newSource)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSource)
}
