package main

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	w := singleLinkList[int]()
	w.Add(1)
	expect := "1"

	if w.ToString() != expect {
		t.Errorf("Expected: %v, found: %v", expect, w.ToString())
	}
}

func TestAddEmpty(t *testing.T) {
	w := singleLinkList[string]()
	w.Add("")
	expect := ""

	if w.ToString() != expect {
		t.Errorf("Expected: %v, found: %v", expect, w.ToString())
	}
}

func TestInsertInt(t *testing.T) {
	w := singleLinkList[int]()
	w.Add(1)
	w.Add(2)
	w.Add(4)
	w.Insert(3, 2)
	expect := "4, 3, 2, 1"

	if w.ToString() != expect {
		t.Errorf("Expect: %v, Found: %v", expect, w.ToString())
	}
}

func TestInsertErr(t *testing.T) {
	w := singleLinkList[int]()
	w.Add(1)
	w.Add(2)
	w.Add(4)
	w.Insert(0, 0)
	expect := "0"

	if w.ToString() != expect {
		t.Errorf("Expect: %v, Found: %v", expect, w.ToString())
	}
}
