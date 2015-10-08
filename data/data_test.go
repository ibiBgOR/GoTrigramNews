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

package data

import (
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	createDatabase("root", "")
}

func TestTitle(t *testing.T) {
	cases := []string{
		"Hello world",
		"This is a News Title",
		"",
		"»",
	}
	for _, c := range cases {
		newsID := postNews(c)
		got := getNewsTitle(newsID)
		if got != c {
			t.Errorf("save and read title: \n saved:  %q,\n got:     %q", c, got)
		}
	}
}

func TestTrigrams(t *testing.T) {
	cases := []struct {
		id       int
		trigrams []string
	}{
		{1, []string{"Hel", "ell", "llo"}},
		{2, []string{"Tes", "est"}},
		{1, []string{"Wie", "ied", "ede", "der"}}, //duplicate id -> no problem (?)
		// {1, []string{"Hel", "ell", "llo"}}, TODO: better error handling (here: duplicate entry)
	}
	for _, c := range cases {
		for _, tri := range c.trigrams {
			putTrigram(tri, c.id)
		}

		for _, tri := range c.trigrams {
			ids := getIdsOfTrigram(tri)
			contains := false
			for _, id := range ids {
				if id == c.id {
					contains = true
				}
			}
			if !contains {
				t.Errorf("%q with id %q is not in %q", tri, string(c.id), ids)
			}
		}

	}
}
