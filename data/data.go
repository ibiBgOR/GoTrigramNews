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

import (
	"github.com/gchaincl/dotsql"
	"github.com/gchaincl/gotic/fs"
)

// getIds(trigram string) ids []int -> returns ids with this trigram
func getIds(trigram string) []int {
	return []int{}
}

// getNewsTitle(id int) string -> returns news title with this id
func getNewsTitle(id int) string {
	return ""
}

// putTrigram(trigram string, id int) -> saves trigram with this id
func putTrigram(trigram string, id int) {
	return
}

// postNews(title string) id int -> saves a new news title, returns id
func postNews(title string) int {
	return 0
}
