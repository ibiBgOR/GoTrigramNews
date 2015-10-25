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
	"testing"
	"strconv"
)

func TestVectorMultiplication (t *testing.T) {

	var vector_1 []int = []int{10, 5}
	var vector_2 []int = []int{2, 3}

	result := vectorMultiplication(vector_1, vector_2)

	if result != 35 {
		t.Log("Result expected to be '35' but was '" + strconv.FormatFloat(result, 'f', 6, 64) + "'")
		t.Fail()
	}
}

func TestVectorNorm (t *testing.T) {

	var vector []int = []int{3, 4}

	result := vectorNorm(vector)

	if result != 5 {
		t.Log("Result expected to be '5' but was '" + strconv.FormatFloat(result, 'f', 6, 64) + "'")
		t.Fail()
	}

}

func TestCosineSimilarity_Similar (t *testing.T) {

	var vector_1 []int = []int{5, 7, 2010, 14, 45, 10, 0, 0, 2}
	var vector_2 []int = []int{23, 12, 2009, 9, 9, 6, 43, 178, 2}

	result := CosineSimilarity(vector_1, vector_2)
	rounded_solution := 0.995664738306

	if 1 - result > 1 - rounded_solution + 0.0000000001 { // We need to check the range, because of infinite resolution of the result
		t.Log("Result expected to be near '" + strconv.FormatFloat(rounded_solution, 'f', 12, 64) + "' but was " + strconv.FormatFloat(result, 'f', 12, 64) + "'")
		t.Fail()
	}

}

func TestCosineSimilarity_NotSimilar (t *testing.T) {

	var vector_1 []int = []int{1, 0, 0, 0, 0}
	var vector_2 []int = []int{0, 1, 1, 1, 1}

	result := CosineSimilarity(vector_1, vector_2)


	if result != 0.0 { // The result is 0, because the two vectors are not similar at all
		t.Log("Result expected to be exact '0.00000' but was " + strconv.FormatFloat(result, 'f', 6, 64) + "'")
		t.Fail()
	}

}

func TestNormalizeTwoVectors(t *testing.T) {
	v_1 := []Vector_element {
		Vector_element{1, "Hel"},
		Vector_element{1, "ell"},
		Vector_element{1, "llo"},
	}
	v_2 := []Vector_element {
		Vector_element{1, "Hel"},
		Vector_element{1, "ell"},
		Vector_element{1, "llo"},
	}
	/*v_2 := []element {
		element{1, "Hel"},
		element{1, "ell"},
		element{1, "llo"},
		element{1, "lo "},
		element{1, "o W"},
		element{1, " Wo"},
		element{1, "Wor"},
		element{1, "orl"},
		element{1, "rld"},
	}*/

	vec_1, vec_2 := NormalizeTwoVectors(v_1, v_2)

	for count := range(vec_1) {
		if vec_1[count] != vec_2[count] {
			t.Logf("The two result vectors are not the same. Vector 1: '%v' Vector 2: '%v'", vec_1, vec_2)
			t.Fail()
		}
	}
}
