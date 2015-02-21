// fetch is a wrapper around package adn
//
// Gets Global feed from ADN (App.Net) and prints username and post text
package main

import (
	//	"fetch/adn"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	db, err := bolt.Open("blog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return b.Put([]byte("0xMF"), []byte("Hey, it worked!"))
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte("0xMF"))
		fmt.Printf("that was: %s\n", v)
		return nil
	})

	/*
		r, err := adn.GetGlobal()
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range r.Data {
			fmt.Printf("%s: %s\n", p.User.UserName, p.Text)
		}
	*/
}
