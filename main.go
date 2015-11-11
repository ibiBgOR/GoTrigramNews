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
	"flag"
	"fmt"
	"github.com/ibiBgOR/GoTrimapNews/ai"
	"github.com/ibiBgOR/GoTrimapNews/bestmatches"
	"github.com/ibiBgOR/GoTrimapNews/data"
)

var database_name = "trigramnews"

func main() {
	// Definition of Command-Line Params
	param_data_file := flag.String("f", "", "path to a file containing News-titles")
	//	param_interactively := flag.Bool("i", false, "run interactively")
	param_purge := flag.Bool("purge", false, "reinitialize the database, purging all existing data")
	param_cosine := flag.Bool("cosine", false, "calculate CosineSimilarity instead of ngram-clusering")
	param_title := flag.String("title", "", "print out similar news-titles")
	param_count := flag.Int("n", 3, "how many similar news titles should be printed?")

	flag.Parse()

	data.InitializeDatabase("root", "")
	data.Connect(database_name, *param_purge)

	if len(*param_data_file) > 0 {
		data.ParseFile(*param_data_file)
	}

	if !*param_cosine {
		if len(*param_title) > 0 {
			for _, match := range bestmatches.GetBestMatches(*param_title, *param_count) {
				fmt.Println(match)
			}
		}
	}

	if *param_cosine {

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

		fmt.Printf("+-----------------------------+\n"+
			"| Cosine similarity: %f |\n"+
			"+-----------------------------+",
			ai.CosineSimilarity(ai.NormalizeTwoVectors(data_1, data_2)))

	}
}
