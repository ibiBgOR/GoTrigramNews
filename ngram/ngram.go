package ngram

import (
	"strings"
)

func BuildNGram(sequence string, n int) []string {
	var result []string
	for seqPosition := 0; seqPosition+n-1 < len(sequence); seqPosition += 1 {
		var ngram []string
		for ncounter := 0; ncounter < n; ncounter += 1 {
			ngram = append(ngram, string(sequence[seqPosition+ncounter]))
		}
		result = append(result, strings.Join(ngram, ""))
	}
	return result
}
