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
	"github.com/ibiBgOR/GoTrimapNews/ngram"
	"github.com/ibiBgOR/GoTrimapNews/data"
)

var database_name = "ai_news_titles"

func main() {
	// 1. Retrieve the filename
	var file_name string
	fmt.Scanln(&file_name)

	// 2. Read all data from file
	var content string = data.ReadFile(file_name)

	// 3. Extract data from file to news titles
	var news_lines []string = data.ExtractNewsLine(content)

	// 4. Save all to the database
	data.InitializeDatabase("root", "")
	data.Connect(database_name, false)

	// 5. For each news line save all n-grams into the database
	for _, line := range news_lines {
		id := data.PostNews(line)
		for _, trigram := range ngram.BuildNGram(line, 3) {
			data.PutTrigram(trigram, id)
		}
	}


}

func main_read_line() {
	// Read a News Title from stdin
	var input string
	fmt.Scanln(&input)
	// Print Ngrams for given News Title
	fmt.Println(ngram.BuildNGram(input, 3))
}
