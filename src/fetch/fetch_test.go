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

func TestToFromSqlite(t *testing.T) {
	r, _ := adn.GetGlobal()
	res, err := ToFromSqlite(r, "data/blog_test.db")
	if len(res) == 0 || err != nil {
		t.Error("Expected result, got nil")
	}
	if len(r.Data) != len(res) {
		t.Error("Returned ADN results and persistence store returned aren't the same", len(r.Data), len(res))
	}
}

func TestToFromBolt(t *testing.T) {
	r, _ := adn.GetGlobal()
	res, err := ToFromBolt(r, "data/blog_test.db")
	if len(res) == 0 || err != nil {
		t.Error("Expected result, got nil")
	}
	if len(r.Data) != len(res) {
		t.Error("Returned ADN results and persistence store returned aren't the same", len(r.Data), len(res))
	}
}
