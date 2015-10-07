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

package data

import "testing"

func TestCreateDatabase(t *testing.T) {
	createDatabase("root", "")
}

func TestTitle(t *testing.T) {
	cases := []string{
		"Hello world",
		"This is a News Title",
	}
	for _, c := range cases {
		newsID := postNews(c)
		got := getNewsTitle(newsID)
		if got != c {
			t.Errorf("save and read title: \n saved:  %q,\n got:     %q", c, got)
		}
	}
}
