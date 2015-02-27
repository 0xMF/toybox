// fetch is a wrapper around package adn
//
// Gets Global feed from ADN (App.Net) and prints username and post text
package main

import (
	"database/sql"
	"fetch/adn"
	"fmt"
	"github.com/boltdb/bolt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var file = "data/blog.db"

func ToFromSqlite(r adn.Response, file string) ([]string, error) {

	os.Remove(file)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table posts (id integer not null primary key, user text, post text);
	delete from posts;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into posts(id, user, post) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for it, p := range r.Data {
		_, err = stmt.Exec(it, p.User.UserName, p.Text)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, user, post from posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var values []string
	var items = 0
	for rows.Next() {
		var id int
		var user string
		var post string
		rows.Scan(&id, &user, &post)
		values = append(values, string(id)+user+post)
		items++
	}
	rows.Close()
	fmt.Printf("%d items retrieved\n", items)

	return values, err
}

func check_error_status(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ToFrom(r adn.Response, file string) ([]string, error) {
	// create new file (remove if exists)
	os.Remove(file)

	db, err := bolt.Open(file, 0600, nil)
	check_error_status(err)
	defer db.Close()

	// store posts in file
	var bName = "posts"
	var items int
	for it, p := range r.Data {
		db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bName))
			if err != nil {
				return err
			}
			return b.Put([]byte(p.Id), []byte(p.User.UserName+p.Text))
		})
		items = it + 1
	}
	fmt.Printf("%d items stored\n", items)

	var values []string
	items = 0

	// display again
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			values = append(values, string(v[:]))
			items++
		}
		return nil
	})
	fmt.Printf("%d items retrieved\n", items)

	return values, err
}

func main() {
	r, err := adn.GetGlobal()
	check_error_status(err)

	//results, err := ToFrom(r, file)
	results, err := ToFromSqlite(r, file)
	check_error_status(err)

	for _, v := range results {
		fmt.Printf("%s\n", v)
	}
}
