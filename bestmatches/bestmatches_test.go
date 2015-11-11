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
	"testing"
)

func TestSort(t *testing.T) {
	matches := []int{1, 1, 2, 3, 3, 3, 3, 4, 4}
	t.Log("unsorted: ", matches)

	frequencies := map[int]int{
		1: 2,
		2: 1,
		3: 4,
		4: 4,
	}
	sorted := SortByFrequency(matches, frequencies)
	t.Log("sorted:   ", sorted)
	RemoveDuplicates(&sorted)
	t.Log("finally:  ", sorted)
	t.Log(frequencies)
}
