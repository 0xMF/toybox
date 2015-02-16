package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	type Marker struct {
		id         string
		name       string
		Percentage int
		Updated_at string
		Version    string
	}
	type Meta struct {
		code   int
		marker Marker
		max_id string
		min_id string
		more   bool
	}
	type Response struct {
		data string
		meta Meta
	}

	adn_global := "https://api.app.net/posts/stream/global"

	res, err := http.Get(adn_global)
	if err != nil {
		log.Fatal(err)
	}
	//global, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	dec := json.NewDecoder(io.Reader(res.Body))
	for {
		var r Response
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
    fmt.Printf("%+v\n", r)
	}
}
