// fetch is a wrapper around package adn
//
// Gets Global feed from ADN (App.Net) and prints username and post text
package main

import (
	"fetch/adn"
	"fmt"
	"log"
)

func main() {
	r, err := adn.GetGlobal()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range r.Data {
		fmt.Printf("%s: %s\n", p.User.UserName, p.Text)
	}
}
