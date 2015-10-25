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
	"github.com/gchaincl/dotsql"
	//	"github.com/gchaincl/gotic/fs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"github.com/ibiBgOR/GoTrimapNews/ai"
)

var sqlConnStr string = ""
var databaseName string = "trigramnews"

// We save the SQL Statements in a map to execute by name and handle the exception in one place!
type stmtFunc func(string) sql.Result

var database *sql.DB

var statementsFile *dotsql.DotSql;

func InitializeDatabase(user string, passwd string) {
	sqlConnStr = user
	if passwd != "" {
		sqlConnStr += ":" + passwd
	}
	sqlConnStr += "@/"

	loadStatements()
}

func loadStatements() {
	dot, err := dotsql.LoadFromFile("data/queries.sql")
	if err != nil {
		panic(err)
	}
	statementsFile = dot
}

func runSql(sql string, args ...interface{}) {
	_, err := statementsFile.Exec(database, sql, args...)
	if err != nil {
		panic(err)
	}
}

func querySql(sql string, args ...interface{}) *sql.Rows {
	result, err := statementsFile.Query(database, sql, args...)
	if err != nil {
		panic(err)
	}
	return result
}

// drop and creates a new database
func createDatabase() {
	if sqlConnStr == "" {
		panic("The database was not initialized. Call 'InitializeDatabase' first and then try again.")
	}
	db, err := sql.Open("mysql", sqlConnStr)
	if err != nil {
		panic(err)
	}

	database = db

	runSql("drop-database")
	runSql("create-database")
	runSql("use-database")
	runSql("create-titles-table")
	runSql("create-trigrams-table")
}

func Connect(databaseName string, create bool) {
	if create {
		createDatabase()
	} else {
		db, err := sql.Open("mysql", sqlConnStr + databaseName)
		if err != nil {
			panic(err)
		}
		database = db
	}

}

// returns all news-ids of a trigram
func GetIdsOfTrigram(trigram string) []int {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	rows := querySql("select-titleids-from-trigram", trigram)

	var ids []int
	for rows.Next() {
		var id_string string
		if err := rows.Scan(&id_string); err != nil {
			panic(err)
		}
		id, err := strconv.Atoi(id_string)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}
	return ids
}

// returns news-title with this id
func GetNewsTitle(id int) string {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	rows := querySql("select-title", id)
	var title string
	for rows.Next() {
		if err := rows.Scan(&title); err != nil {
			panic(err)
		}
	}
	return title
}

// saves trigram with this id
func PutTrigram(trigram string, id int) {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	runSql("insert-trigram", id, trigram)
}

// saves a new news title, returns id
func PostNews(title string) int {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	runSql("insert-title", title)
	rows := querySql("select-titleid-by-name", title)

	var id_string string
	for rows.Next() {
		if err := rows.Scan(&id_string); err != nil {
			panic(err)
		}
	}
	id, err := strconv.Atoi(id_string)
	if err != nil {
		panic(err)
	}
	return id
}

func GetTrigramsByTitle(title string) []ai.Vector_element {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	rows := querySql("select-all-trigrams-by-title", title)

	var result []ai.Vector_element;
	for rows.Next() {
		var nextElementName string;
		var nextElementCount int;
		if err := rows.Scan(&nextElementName, &nextElementCount); err != nil {
			panic(err)
		}
		elem := ai.Vector_element{
			Count: nextElementCount,
			Ngram: nextElementName,
		}
		result = append(result, elem)
	}

	return result
}

func GetCountOfTitles() int {

	if sqlConnStr == "" {
		panic("Database not initialized")
	}

	rows := querySql("count-all-titles")

	for rows.Next() {
		var nextElement int;
		if err := rows.Scan(&nextElement); err != nil {
			panic(err)
		}
		return nextElement
	}

	return 0;
}
