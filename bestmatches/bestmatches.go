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

package bestmatches

import (
	"github.com/ibiBgOR/GoTrimapNews/data"
	"github.com/ibiBgOR/GoTrimapNews/ngram"
	"sort"
)

// returns n=count titles matching to the title
func GetBestMatches(title string, count int) []string {
	var trigram_matches []int

	// first calculate ngrams of the search string
	for trigram := range ngram.BuildNGram(title, 3) {
		for match := range data.GetIdsOfTrigram(trigram) {
			trigram_matches = append(trigram_matches, match)
		}
	}

	// now get the *count* most frequent news ids
	// TODO: sort by occurence and remove duplicates
	sort.Ints(trigram_matches)
	var bestMatches []int
	for i := 0; i < count; i += 1 {
		bestMatches = append(bestMatches, trigram_matches[i])
	}

	// get the according titles
	var titles []string
	for id := range bestMatches {
		titles = append(titles, data.GetNewsTitle(id))
	}

	return titles
}
