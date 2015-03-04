// Copyright 2015, Mark Fernandes. All rights reserved.
// Use of this source code is governed by a ISC-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fetch/adn"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var file = "data/blog.db"

func check_error_status(err error, args ...string) {
	if err != nil {
		log.Fatal(err)
	}
	if len(args) > 0 {
		log.Printf("%q: %s\n", err, args)
	}
}

func ToFromSqlite(r adn.Response, file string) ([]string, error) {

	os.Remove(file)
	db, err := sql.Open("sqlite3", file)
	check_error_status(err)
	defer db.Close()

	sqlStmt := `
	create table posts (id integer not null primary key, user text, post text);
	delete from posts;
	`
	_, err = db.Exec(sqlStmt)
	check_error_status(err, sqlStmt)

	tx, err := db.Begin()
	check_error_status(err)

	stmt, err := tx.Prepare("insert into posts(id, user, post) values(?, ?, ?)")
	check_error_status(err)
	defer stmt.Close()

	for it, p := range r.Data {
		_, err = stmt.Exec(it, p.User.UserName, p.Text)
		check_error_status(err)
	}
	tx.Commit()

	rows, err := db.Query("select id, user, post from posts")
	check_error_status(err)
	defer rows.Close()
	var values []string
	var items = 0
	for rows.Next() {
		var id int
		var user string
		var post string
		rows.Scan(&id, &user, &post)
		fmt.Println(id, user, post)
		values = append(values, "@"+user+": "+post)
		items++
	}
	rows.Close()
	fmt.Printf("%d items retrieved\n", items)

	return values, err
}

func main() {

	rchan := make(chan adn.Response)

	go func() {
		r, err := adn.GetGlobal()
		check_error_status(err)
		rchan <- r
	}()

	go func() {
		r := <-rchan
		results, err := ToFromSqlite(r, file)
		check_error_status(err)

		for _, v := range results {
			fmt.Printf("%s\n", v)
		}
	}()

	fmt.Printf("Press any key to continue...")
	var input string
	fmt.Scanln(&input)
}
