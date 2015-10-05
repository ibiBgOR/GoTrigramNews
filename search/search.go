package main

import (
	"GoTrimapNews/ngram"
	"fmt"
)

func main() {
	fmt.Println(ngram.BuildNGram("Hello World", 3))
}
