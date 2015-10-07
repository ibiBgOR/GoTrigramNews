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

package ngram

import "testing"

func TestNgram(t *testing.T) {
	cases := []struct {
		n    int
		str  string
		want []string
	}{
		{3, "Hello world", []string{"Hel", "ell", "llo", "lo ", "o w", " wo", "wor", "orl", "rld"}},
		// How should strings shorter than n be handled?
		{3, "He", []string{}},
		// How should empty strings be handled?
		{1, "", []string{}},
		{2, "Hello world", []string{"He", "el", "ll", "lo", "o ", " w", "wo", "or", "rl", "ld"}},
	}

	for _, c := range cases {
		err := false
		got := BuildNGram(c.str, c.n)
		if len(got) == len(c.want) {
			for index, _ := range got {
				if got[index] != c.want[index] {
					err = true
				}
			}
		} else {
			err = true
		}
		if err {
			t.Errorf("BuildNGram(%q, %q): \n got:  %q,\n want: %q", c.str, string(c.n), got, c.want)
		}
	}
}
