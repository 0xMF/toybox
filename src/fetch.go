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

	type HashTag struct {
		len  int
		name string
		pos  int
	}
	type Link struct {
		len  int
		pos  int
		text string
		url  string
	}
	type Mention struct{}
	type Entity struct {
		hashtags []HashTag
		links    []Link
		mentions []Mention
	}
	type Source struct {
		client_id string
		link      string
		name      string
	}
	type User struct {
	}
	type Location struct {
		latitude  float32
		longitude float32
	}
	type Annotation struct {
		atype string // change later
		value Location
	}

	type Post struct {
		canonical_url string
		created_at    string
		entities      Entity
		html          string
		id            string
		machine_only  bool
		num_replies   int
		num_reposts   int
		num_stars     int
		source        Source
		text          string
		thread_id     string
		user          User
		you_reposted  bool
		you_starred   bool
		annotations   []Annotation
		reposters     []User
		starred_by    []User
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
