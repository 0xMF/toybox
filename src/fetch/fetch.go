// fetch is a wrapper around package adn
//
// Gets Global feed from ADN (App.Net) and prints username and post text
package main

import (
	"fetch/adn"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

var file = "data/blog.db"

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

	results, err := ToFrom(r, file)
	check_error_status(err)

	for _, v := range results {
		fmt.Printf("%s\n", v)
	}
}
