/*
Copyright 2015 Oscar Ruckdeschel, Janik Schmidt, Jonathan Kuhse.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"github.com/ibiBgOR/GoTrimapNews/ngram"
	"github.com/ibiBgOR/GoTrimapNews/data"
	"github.com/ibiBgOR/GoTrimapNews/ai"
)

var database_name = "trigramnews"

func main() {
	// 1. Retrieve the filename
	var file_name string
	fmt.Print("Enter filename: ")
	fmt.Scanln(&file_name)

	// 2. Read all data from file
	log.Println("Start reading the file.")
	var content string = data.ReadFile(file_name)
	log.Println("File was read.")

	// 3. Extract data from file to news titles
	log.Println("Start extracting the data from the file.")
	var news_lines []string = data.ExtractNewsLine(content)
	log.Println("Data was extracted.")

	// 4. Save all to the database
	log.Println("Initialize database.")
	data.InitializeDatabase("root", "")
	data.Connect(database_name, true)
	log.Println("Database initialized.")

	// 5. For each news line save all n-grams into the database
	log.Println("Moving all data to the database.")
	for _, line := range news_lines {
		id := data.PostNews(line)
		for _, trigram := range ngram.BuildNGram(line, 3) {
			data.PutTrigram(trigram, id)
		}
	}
	log.Println("Data saved.")

	titleCount := data.GetCountOfTitles()

	// 6. Prepare to calculate distances
	var title_1 int
	var title_2 int

	fmt.Printf("Enter number for first title (%d): ", titleCount)
	fmt.Scanln(&title_1)
	fmt.Printf("Enter number for second title (%d): ", titleCount)
	fmt.Scanln(&title_2)

	data_1 := data.GetTrigramsByTitle(data.GetNewsTitle(title_1))
	data_2 := data.GetTrigramsByTitle(data.GetNewsTitle(title_2))

	fmt.Printf("+-----------------------------+\n" +
	           "| Cosine similarity: %f |\n" +
			   "+-----------------------------+",
		ai.CosineSimilarity(ai.NormalizeTwoVectors(data_1, data_2)))
}

func main_read_line() {
	// Read a News Title from stdin
	var input string
	fmt.Scanln(&input)
	// Print Ngrams for given News Title
	fmt.Println(ngram.BuildNGram(input, 3))
}
