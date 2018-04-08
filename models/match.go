package models

type Match struct {
	Score float64
	Value string
}

func NewCommandMatch(
	score float64,
	value string,
	commands []*Command,
	CallDirectory string,
) Match {

	score *= float64(len(commands))
	return Match{
		Score: score,
		Value: value,
	}
}

func NewDirectoryMatch(
	score float64,
	value string,
	commands []*Command,
) Match {

	score *= float64(len(commands))
	return Match{
		Score: score,
		Value: value,
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
	return m[i].Score < m[j].Score
}
