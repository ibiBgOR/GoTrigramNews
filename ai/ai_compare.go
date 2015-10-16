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

func CosineSimilarity (vector_1 []int, vector_2 []int) float64 {
	return vectorMultiplication(vector_1, vector_2) / (vectorNorm(vector_1) * vectorNorm(vector_2))
}

func vectorMultiplication (vector_1 []int, vector_2 []int) float64 {
	if len(vector_1) != len(vector_2) {
		panic("The vectors does not have the same length.")
	}
	var result float64 = 0

	for i, elem_1 := range vector_1 {
		result += float64(elem_1 * vector_2[i])
	}

	return result
}

func vectorNorm (vector []int) float64 {
	var result float64 = 0

	for _, elem_1 := range vector {
		result += float64(elem_1 * elem_1)
	}

	return math.Sqrt(result)
}