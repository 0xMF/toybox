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

func main() {

	// Get Global posts
	r, err := adn.GetGlobal()
	if err != nil {
		log.Fatal(err)
	}

	// create new file (remove if exists)
	var file = "/tmp/blog.db"
	var bName = "posts"
	os.Remove(file)

	db, err := bolt.Open(file, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// store posts in file
	for _, p := range r.Data {
		db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bName))
			if err != nil {
				return err
			}
			return b.Put([]byte(p.Created_At), []byte(p.User.UserName+p.Text))
		})
	}

	// display again
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s\n", v)
		}
		return nil
	})
}
