package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toddsifleet/godo/utils"
)

func TestMatchWord(t *testing.T) {
	tests := []struct {
		search []rune
		data   []rune
		start  int
		end    int
	}{
		{[]rune("foo"), []rune("foobar"), 0, 2},
		{[]rune("foo"), []rune("fozobar"), 0, 3},
		{[]rune("foo"), []rune("ofozobar"), 1, 4},
	}

	for _, test := range tests {
		start, end := utils.MatchWord(test.search, test.data)
		assert.Equal(t, test.start, start, "should have the correct start")
		assert.Equal(t, test.end, end, "should have the correct end")
	}
}

func TestScoreMatch(t *testing.T) {
	tests := []struct {
		start     int
		end       int
		lenSearch int
		lenData   int
		score     float64
	}{
		{0, 2, 3, 6, 1},
		{0, 2, 3, 3, 1},
		{0, 4, 3, 6, .8},
		{0, 0, 3, 6, 0},
	}

	for _, test := range tests {
		score := utils.ScoreMatch(test.start, test.end, test.lenSearch, test.lenData)
		assert.Equal(t, test.score, score, "should have the correct score")
	}
}

func TestCompareScoreMatch(t *testing.T) {
	tests := []struct {
		search []rune
		good   []rune
		bad    []rune
	}{
		{[]rune("foo"), []rune("foobar"), []rune("fozobar")},
		{[]rune("foo"), []rune("foobar"), []rune("ffoobar")},
	}

	for _, test := range tests {
		start, end := utils.MatchWord(test.search, test.good)
		scoreGood := utils.ScoreMatch(start, end, len(test.search), len(test.good))
		start, end = utils.MatchWord(test.search, test.bad)
		scoreBad := utils.ScoreMatch(start, end, len(test.search), len(test.bad))
		assert.True(
			t,
			scoreGood > scoreBad,
			fmt.Sprintf("%f !> %f", scoreGood, scoreBad),
		)
	}
}

func TestScoreWords(t *testing.T) {
	tests := []struct {
		words [][]rune
		data  []rune
		score float64
	}{
		{[][]rune{[]rune("foo"), []rune("bar")}, []rune("foobar"), 1.0},
	}

	for _, test := range tests {
		score := utils.ScoreWords(test.words, test.data)
		assert.Equal(t, test.score, score, "should have the correct score")
	}
}
