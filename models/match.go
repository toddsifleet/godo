package models

import "time"

type Match struct {
	Score float64
	Value string
	Times []*time.Time
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
