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

package ai

import (
	"math"
)

type Vector_element struct {
	Count int
	Ngram string
}

// We want to return two int-vectors which we can compare
// Therefor we need to construct the vectors out of the given ones
func NormalizeTwoVectors(vector_1 []Vector_element, vector_2 []Vector_element) ([]int, []int) {
	// We check all containing ngrams (how often do they appear, do they appear in the other vector, ...)
	var vector_1_new = vector_1
	var vector_2_new = vector_2

	var vec_1 []int
	var vec_2 []int
	lastElement := Vector_element{-1, ""}

	for counter_1, elem_1 := range vector_1 {
		if lastElement.Count != -1 && lastElement == elem_1 {
			break
		}
		lastElement = elem_1

		contains := false
		for counter_2, elem_2 := range vector_2 {
			if elem_1.Ngram == elem_2.Ngram {
				// if both vectors contains the same ngam: add it to the result vectors
				contains = true
				vec_1 = append(vec_1, elem_1.Count)
				vec_2 = append(vec_2, elem_2.Count)

				// we also delete the ngram from both, so we don't add it twice!
				if len(vector_1_new) != counter_1 {
					vector_1_new = vector_1_new[:counter_1+copy(vector_1_new[counter_1:], vector_1_new[counter_1+1:])]
				}
				if len(vector_2_new) != counter_2 {
					vector_2_new = vector_2_new[:counter_2+copy(vector_2_new[counter_2:], vector_2_new[counter_2+1:])]
				}

				break
			}
		}
		if !contains {
			// if not both vectors contains the same ngram: add to the first vector the count and '0' to the second
			vec_1 = append(vec_1, elem_1.Count)
			vec_2 = append(vec_2, 0)
		}
	}

	// we have to iterate over the elements in the second vector for the unique elements in this vector
	for _, elem_2 := range vector_2_new {
		contains := false
		for _, elem_1 := range vector_1_new {
			if elem_1.Ngram == elem_2.Ngram {
				// if both vectors contains the same ngam: add it to the result vectors
				contains = true
				vec_1 = append(vec_1, elem_1.Count)
				vec_2 = append(vec_2, elem_2.Count)
				break
			}
		}
		if !contains {
			// if not both vectors contains the same ngram: add to the first vector the count and '0' to the second
			vec_1 = append(vec_1, 0)
			vec_2 = append(vec_2, elem_2.Count)
		}
	}

	return vec_1, vec_2
}

func CosineSimilarity(vector_1 []int, vector_2 []int) float64 {
	return vectorMultiplication(vector_1, vector_2) / (vectorNorm(vector_1) * vectorNorm(vector_2))
}

func vectorMultiplication(vector_1 []int, vector_2 []int) float64 {
	if len(vector_1) != len(vector_2) {
		panic("The vectors does not have the same length.")
	}
	var result float64 = 0

	for i, elem_1 := range vector_1 {
		result += float64(elem_1 * vector_2[i])
	}

	return result
}

func vectorNorm(vector []int) float64 {
	var result float64 = 0

	for _, elem_1 := range vector {
		result += float64(elem_1 * elem_1)
	}

	return math.Sqrt(result)
}
