package main

import (
  "fetch/adn"
	"testing"
)


func TestGlobalData(t *testing.T) {
	r, _ := adn.GetGlobal()
	if r.Meta.Code != 200 {
		t.Error("Expected 200, got ", r.Meta.Code)
	}
}

func TestToFrom(t *testing.T) {
	r, _ := adn.GetGlobal()
  res, err:=ToFrom(r,"data/blog_test.db")
  if len(res) == 0 || err != nil {
		t.Error("Expected result, got ni")
  }
  if len(r.Data) != len(res) {
    t.Error("Expected got from ADN and persistence store", len(r.Data), len(res) )
  }
}
