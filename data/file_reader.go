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
	"io/ioutil"
	"strings"
)

func ReadFile(fileName string) string {
	dataArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(dataArray)
}

func ExtractNewsLine(content string) []string {
	result := strings.Split(content, "\n")

	for count, elem := range result {
		result[count] = elem[strings.IndexAny(elem, "\t") + 1:]
	}

	return result
}
