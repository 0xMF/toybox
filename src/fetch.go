package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	type HashTag struct {
		Len  int    `json:"len"`
		Name string `json:"name"`
		Pos  int    `json:"pos"`
	}
	type Link struct {
		Url  string `json:"url"`
		Text string `json:"text"`
		Pos  int    `json:"pos"`
		Len  int    `json:"len"`
	}
	type Mention struct{}
	type Entity struct {
		Mentions []Mention `json:"mentions"`
		Hashtags []HashTag `json:"hashtags"`
		Links    []Link    `json:"links"`
	}
	type Source struct {
		Link      string `json:"link"`
		Name      string `json:"name"`
		Client_Id string `json:"client_id"`
	}
	type Image struct {
		Url        string `json:"url"`
		Width      int    `json:"width"`
		Is_Defualt bool   `json:"is_default"`
		Height     int    `json:"height"`
	}
	type Description struct {
		Text     string `json:"text"`
		Html     string `json:"html"`
		Entities Entity `json:"entity"`
	}
	type Counts struct {
		Following int `json:"following"`
		Posts     int `json:"posts"`
		Stars     int `json:"stars"`
	}
	type User struct {
		UserName      string      `json:"username"`
		Avatar        Image       `json:"avatar"`
		Description   Description `json:"description"`
		Locale        string      `json:"locale"`
		Created_At    string      `json:"created_at"`
		Canonical_Url string      `json:"canonical_url"`
		Cover         Image       `json:"cover_image"`
		Timezone      string      `json:"timezone"`
		Counts        Counts      `json:"counts"`
		Type          string      `json:"type"`
		Id            string      `json:"id"`
		Name          string      `json:"name"`
	}
	type Value struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	}
	type Annotation struct {
		Type  string `json:"type"`
		Value Value  `json:"value"`
	}

	type Post struct {
		Created_At    string       `json:"created_at"`
		Num_Stars     int          `json:"num_stars"`
		Num_Replies   int          `json:"num_replies"`
		Source        Source       `json:"source"`
		Text          string       `json:"text"`
		Num_Reposts   int          `json:"num_reposts"`
		Id            string       `json:"id"`
		Canonical_Url string       `json:"canonical_url"`
		Entities      Entity       `json:"entities"`
		Html          string       `json:"html"`
		Machine_Only  bool         `json:"machine_only"`
		User          User         `json:"user"`
		Thread_Id     string       `json:"thread_id"`
		Pagination_Id string       `json:"pagination_id"`
		You_Reposted  bool         `json:"you_reposted"`
		You_Starred   bool         `json:"you_starred"`
		Annotations   []Annotation `json:"annotations"`
		Reposters     []User       `json:"reposters"`
		Starred_By    []User       `json:"starred_by"`
	}

	type Marker struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Percentage int    `json:"percentage"`
		Updated_At string `json:"updated_at"`
		Version    string `json:"version"`
	}
	type Meta struct {
		Min_Id string `json:"min_id"`
		Code   int    `json:"code"`
		Marker Marker `json:"marker"`
		Max_Id string `json:"max_id"`
		More   bool   `json:"more"`
	}
	type Response struct {
		Meta Meta   `json:"meta"`
		Data []Post `json:"data"`
	}

	adn_global := "https://api.app.net/posts/stream/global"

	res, err := http.Get(adn_global)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var r Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range r.Data {
		fmt.Printf("%s: %s\n", p.User.UserName, p.Text)
	}
}
