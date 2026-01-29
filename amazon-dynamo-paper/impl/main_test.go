package main

import (
	"testing"
)

func TestKeyToPosition_SameKey_SamePosition(t *testing.T) {
	p1 := KeyToPosition([]byte("hello"))
	p2 := KeyToPosition([]byte("hello"))
	if p1 != p2 {
		t.Errorf("expected same position for same key, got %d and %d", p1, p2)
	}
}

func TestKeyToPosition_DifferentKey_DifferentPosition(t *testing.T) {
	p1 := KeyToPosition([]byte("hello"))
	p2 := KeyToPosition([]byte("world"))
	if p1 == p2 {
		t.Errorf("expected different position for different keys, got %d and %d", p1, p2)
	}
}
