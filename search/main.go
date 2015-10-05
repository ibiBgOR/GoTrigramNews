package main

import (
	"fmt"
	"org.hwr/GoTrimapNews/ngram"
)

func main() {
	fmt.Println(ngram.BuildNGram("Hello World", 3))
}
