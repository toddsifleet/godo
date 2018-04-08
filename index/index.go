package index

import (
	"sort"
	"strings"

	"github.com/toddsifleet/godo/models"
	"github.com/toddsifleet/godo/utils"
)

type Index interface {
	AddCommand(c *models.Command)
	GetDirectories(string, string) []models.Match
	GetCommands(string, string) []models.Match
}

type index struct {
	commands    map[string][]*models.Command
	directories map[string][]*models.Command
}

func New() Index {
	return &index{
		commands:    map[string][]*models.Command{},
		directories: map[string][]*models.Command{},
	}
}

func splitAndRune(input string) [][]rune {
	splits := strings.Split(input, " ")

	result := make([][]rune, len(splits))
	for i, word := range splits {
		result[i] = []rune(word)
	}
	return result
}

func (s *index) AddCommand(c *models.Command) {
	if s.commands[c.Value] == nil {
		s.commands[c.Value] = []*models.Command{c}
	} else {
		s.commands[c.Value] = append(s.commands[c.Value], c)
	}

	if s.directories[c.RunDirectory] == nil {
		s.directories[c.RunDirectory] = []*models.Command{c}
	} else {
		s.directories[c.RunDirectory] = append(s.directories[c.RunDirectory], c)
	}
}

func (s *index) GetDirectories(currentDirectory, search string) []models.Match {
	var matches models.MatchList
	runes := splitAndRune(utils.ReverseString(search))
	for directory, commands := range s.directories {
		if directory == currentDirectory {
			continue
		}
		score := utils.ScoreWords(runes, []rune(utils.ReverseString(directory)))
		if score > 0 {
			matches = append(matches, models.NewDirectoryMatch(score, directory, commands))
		}
	}
	sort.Sort(sort.Reverse(matches))
	return matches
}

func (s *index) GetCommands(currentDirectory, search string) []models.Match {
	var matches models.MatchList
	runes := splitAndRune(search)
	for commandValue, commands := range s.commands {
		score := utils.ScoreWords(runes, []rune(commandValue))
		if score > 0 {
			matches = append(matches, models.NewCommandMatch(score, commandValue, commands, currentDirectory))
		}
	}
	sort.Sort(sort.Reverse(matches))
	return matches
}
