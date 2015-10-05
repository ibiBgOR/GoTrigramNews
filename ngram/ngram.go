/*
Copyright 2015 Oscar Ruckdaeschel, Janik Schmidt, Jonathan Kuhse.

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

package ngram

import (
	"strings"
)

func BuildNGram(sequence string, n int) []string {
	var result []string
	for seqPosition := 0; seqPosition+n-1 < len(sequence); seqPosition += 1 {
		var ngram []string
		for ncounter := 0; ncounter < n; ncounter += 1 {
			ngram = append(ngram, string(sequence[seqPosition+ncounter]))
		}
		result = append(result, strings.Join(ngram, ""))
	}
	return result
}
