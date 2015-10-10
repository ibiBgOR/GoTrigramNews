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
    "strconv"
    "fmt"
)

func TestReadFileNotAvaliable(t *testing.T) {
	defer func() {
        if r := recover(); r != nil {
            t.Log("An error occured since the file cannot be found. But we expected this!")
        }
    }()

	ReadFile("i_am_not.here")
}

func TestReadFileAvailable(t *testing.T) {
	result := len(ReadFile("test.file"))

    if result != 42 { // Note that a new line equals two bytes/characters (\n)
        t.Log("The result has not the expected size of 42 != " + strconv.Itoa(result))
        t.Fail()
    }
}

func TestExtractNews(t *testing.T) {
    expected := []string {"Hello World\r", "this is a test\r", "file"}
    resulted := ExtractNewsLine(ReadFile("test.file"))

    for count, element := range expected {
        if resulted[count] != element {
            fmt.Println("Result is not the same as expected. Expected: '", []byte(element), "'\nbut got: '", []byte(resulted[count]), "'")
            t.Fail()
        }
    }
}