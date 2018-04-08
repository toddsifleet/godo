package models

import (
	"math"
	"time"
)

const (
	STRING_SCORE_THRESHOLD    = .1
	DEFUALT_DECAY_RATE        = 1000.
	SAME_DIRECTORY_DECAY_RATE = 10.
)

type Match struct {
	Value string

	stringScore  float64
	commandScore float64
}

func getCommandScore(commands []*Command, targetDirectory string) float64 {
	now := float64(time.Now().In(time.UTC).Unix())
	result := 0.0
	for _, command := range commands {
		decayRate := DEFUALT_DECAY_RATE
		if command.RunDirectory == targetDirectory {
			decayRate = SAME_DIRECTORY_DECAY_RATE
		}
		result += math.Pow(float64(command.Time.Unix())/now, decayRate)
	}
	return result
}

func NewCommandMatch(
	score float64,
	value string,
	commands []*Command,
	callDirectory string,
) Match {
	return Match{
		stringScore:  score,
		commandScore: getCommandScore(commands, callDirectory),
		Value:        value,
	}
}

func NewDirectoryMatch(
	score float64,
	value string,
	commands []*Command,
) Match {

	return Match{
		stringScore:  score,
		commandScore: getCommandScore(commands, ""),
		Value:        value,
	}
}

type MatchList []Match

func (m MatchList) Len() int {
	return len(m)
}

func (m MatchList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchList) Less(i, j int) bool {
	if math.Abs(m[i].stringScore-m[j].stringScore) < STRING_SCORE_THRESHOLD {
		return m[i].commandScore < m[j].commandScore
	}

	return m[i].stringScore < m[j].stringScore
}
