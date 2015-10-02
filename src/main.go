package main

import (
	"fmt"
)

func main() {
	fmt.Print(buildNGram("Hello World", 3))
}

type trigram struct {
	a, b, c string
}

func buildNGram(sequence string, n int) []trigram {
	var result []trigram

	for i := 0; i + 2 < len(sequence); i += 1 {
		result = append(result, trigram{string(sequence[i]), string(sequence[i + 1]), string(sequence[i + 2])})
	}
	return result
}
