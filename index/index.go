package index

import (
	"sort"
	"strings"
	"time"

	"github.com/toddsifleet/godo/models"
	"github.com/toddsifleet/godo/utils"
)

type Index interface {
	AddCommand(c *models.Command)
	GetDirectories(string, string) []models.Match
	GetCommands(string, string) []models.Match
}

type index struct {
	commands    map[string][]*time.Time
	directories map[string]map[string][]*time.Time
}

func New() Index {
	return &index{
		commands:    map[string][]*time.Time{},
		directories: map[string]map[string][]*time.Time{},
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

func flattenTimes(input map[string][]*time.Time) []*time.Time {
	var result []*time.Time

	for _, times := range input {
		result = append(result, times...)
	}
	return result
}

func (s *index) AddCommand(c *models.Command) {
	if s.commands[c.Value] == nil {
		s.commands[c.Value] = []*time.Time{&c.Time}
	} else {
		s.commands[c.Value] = append(s.commands[c.Value], &c.Time)
	}

	if s.directories[c.RunDirectory] == nil {
		s.directories[c.RunDirectory] = map[string][]*time.Time{
			c.Value: []*time.Time{&c.Time},
		}
	} else if s.directories[c.RunDirectory][c.Value] == nil {
		s.directories[c.RunDirectory][c.Value] = []*time.Time{&c.Time}
	} else {
		s.directories[c.RunDirectory][c.Value] = append(s.directories[c.RunDirectory][c.Value], &c.Time)
	}
}

func (s *index) GetDirectories(currentDirectory, search string) []models.Match {
	var matches models.MatchList
	runes := splitAndRune(utils.ReverseString(search))
	for dir, commands := range s.directories {
		if dir == currentDirectory {
			continue
		}
		score := utils.ScoreWords(runes, []rune(utils.ReverseString(dir)))
		if score > 0 {
			matches = append(matches, models.Match{
				Score: score,
				Value: dir,
				Times: flattenTimes(commands),
			})
		}
	}
	sort.Sort(sort.Reverse(matches))
	return matches
}

func (s *index) GetCommands(currentDirectory, search string) []models.Match {
	// TODO: Prefer matches from same directory
	var matches models.MatchList
	runes := splitAndRune(search)
	for cmd, times := range s.commands {
		score := utils.ScoreWords(runes, []rune(cmd))
		if score > 0 {
			matches = append(matches, models.Match{
				Score: score,
				Value: cmd,
				Times: times,
			})
		}
	}
	sort.Sort(sort.Reverse(matches))
	return matches
}
