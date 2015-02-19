package adn

import (
	"testing"
)

func TestGlobal(t *testing.T) {
	r, _ := GetGlobal()
	if r.Meta.Code != 200 {
		t.Error("Expected 200, got ", r.Meta.Code)
	}
}
