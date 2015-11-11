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
	for _, trigram := range ngram.BuildNGram(title, 3) {
		for _, match := range data.GetIdsOfTrigram(trigram) {
			trigram_matches = append(trigram_matches, match)
		}
	}

	// TODO: now get the *count* most frequent news ids
	frequencies := map[int]int{}

	for _, id := range trigram_matches {
		_, contains := frequencies[id]
		if contains {
			frequencies[id] += 1
		} else {
			frequencies[id] = 1
		}
	}

	// now sort according to the count
	bestMatches := (SortByFrequency(trigram_matches, frequencies))
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

type IdFrequency struct {
	Id        int
	Frequency int
}

type ByFrequency []IdFrequency

func (nf ByFrequency) Len() int      { return len(nf) }
func (nf ByFrequency) Swap(i, j int) { nf[i], nf[j] = nf[j], nf[i] }
func (nf ByFrequency) Less(i, j int) bool {
	less := nf[i].Frequency > nf[j].Frequency
	if nf[i].Frequency == nf[j].Frequency {
		less = nf[i].Id < nf[j].Id
	}
	return less
}

func SortByFrequency(ids []int, frequencies map[int]int) []int {
	nf := make(ByFrequency, len(ids))
	for i, id := range ids {
		nf[i] = IdFrequency{id, frequencies[id]}
	}
	sort.Sort(ByFrequency(nf))
	sortedIds := make([]int, len(ids))
	for i, nf := range nf {
		sortedIds[i] = nf.Id
	}
	return sortedIds
}

func RemoveDuplicates(xs *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
