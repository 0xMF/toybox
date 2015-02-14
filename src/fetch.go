package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
  "log"
)

func main() {
	adn_global := "https://api.app.net/posts/stream/global"

	res, err := http.Get(adn_global)
	if err != nil {
		log.Fatal(err)
	}
  global, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",global)
}
