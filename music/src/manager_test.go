package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("New failed")
	}
	if mm := Len() != 0 {
		t.Error("failed")
	}	
	m0 := &MussicEntry {

	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error()
	}

	
}