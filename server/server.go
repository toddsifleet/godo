package server

import (
	"github.com/toddsifleet/godo/index"
	"github.com/toddsifleet/godo/models"
)

type Server struct {
	index index.Index
}

func New(index index.Index) (*Server, error) {
	return &Server{index: index}, nil
}

func matchesToResponse(matches models.MatchList) models.Response {
	response := models.Response{
		Options: make([]string, len(matches)),
	}

	for i, match := range matches {
		response.Options[i] = match.Value
	}
	return response
}

func (s *Server) FindDirectories(request *models.Request, response *models.Response) error {
	matches := s.index.GetDirectories(request.CurrentDirectory, request.SearchTerm)
	*response = matchesToResponse(matches)
	return nil
}

func (s *Server) FindCommands(request *models.Request, response *models.Response) error {
	matches := s.index.GetCommands(request.CurrentDirectory, request.SearchTerm)
	*response = matchesToResponse(matches)
	return nil
}
