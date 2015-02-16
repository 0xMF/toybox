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
		Len  int `json:"len"`
		Name string
		Pos  int
	}
	type Link struct {
		Url  string
		Text string
		Pos  int
		Len  int `json:"len"`
	}
	type Mention struct{}
	type Entity struct {
		Mentions []Mention
		Hashtags []HashTag
		Links    []Link
	}
	type Source struct {
		Link      string
		Name      string
		Client_Id string
	}
	type Image struct {
		Url        string
		Width      int
		Is_Defualt bool
		Height     int
	}
	type Description struct {
		Text     string
		Html     string
		Entities Entity
	}
	type Count struct {
		Following int
		Posts     int
		Stars     int
	}
	type User struct {
		UserName      string      `json:"username"`
		Avatar        Image       `json:"avatar"`
		Profile       Description `json:"description"`
		Locale        string
		Created_At    string
		Canonical_Url string
		Cover         Image `json:"cover_image"`
		Timezone      string
		Counts        Count
		Type          string `json:"type"`
		Id            string
		FullName      string `json:"name"`
	}
	type Location struct {
		Latitude  float32
		Longitude float32
	}
	type Annotation struct {
		Annotation_type string `json:"type"`
		Value           Location
	}

	type Post struct {
		Created_At    string
		Num_Stars     int
		Num_Replies   int
		Source        Source
		Text          string
		Num_Reposts   int
		Id            string
		Canonical_Url string
		Entities      Entity
		Html          string
		Machine_Only  bool
		Post_User     User `json:"user"`
		Thread_Id     string
		Pagination_Id string
		You_Reposted  bool
		You_Starred   bool
		Annotations   []Annotation
		Reposters     []User
		Starred_By    []User
	}

	type Marker struct {
		Id         string
		Name       string
		Percentage int
		Updated_at string
		Version    string
	}
	type Meta struct {
		Min_Id      string `json:"min_id"`
		Code        int    `json:"code"`
		Meta_Marker Marker
		Max_Id      string `json:"max_id"`
		More        bool   `json:"more"`
	}
	type Response struct {
		Meta Meta
		Data []Post
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
		fmt.Printf("%s: %s\n", p.Post_User.UserName, p.Text)
	}
}
