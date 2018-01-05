package utils

const MATCH_COEFFICIENT = float64(1)
const LOCATION_COEFFICIENT = float64(1)

// MatchWord --  Fuzzy match searchTerm with data, and returns the start and end of the match
func MatchWord(searchTerm []rune, data []rune) (int, int) {
	start, current := -1, -1
	for _, l := range searchTerm {
		found := false
		for current < len(data)-1 {
			current++
			if l == data[current] {
				found = true
				if start < 0 {
					start = current
				}
				break
			}
		}
		if !found {
			return 0, 0
		}
	}

	return start, current
}

// ScoreMatch -- Return a value 0-1 (0 is not a match 1 is a perfect match)
func ScoreMatch(start, end, lenSearch, lenData int) float64 {
	if lenSearch == 0 {
		return 1
	}
	if start == 0 && end == 0 {
		return 0
	}
	matchScoreWord := MATCH_COEFFICIENT * float64(lenSearch) / float64(end-start+1)
	locationScoreWord := LOCATION_COEFFICIENT * float64(lenData-start) / float64(lenData)

	return (matchScoreWord + locationScoreWord) / (MATCH_COEFFICIENT + LOCATION_COEFFICIENT)
}

func ScoreWords(searchTerms [][]rune, data []rune) float64 {
	sum := float64(0)
	start, end := -1, -1
	for _, term := range searchTerms {
		lenData := len(data) - start
		start, end = MatchWord(term, data[end+1:])
		if start == 0 && end == 0 {
			return 0
		}
		sum += ScoreMatch(start, end, len(term), lenData)
	}
	return sum / float64(len(searchTerms))
}
