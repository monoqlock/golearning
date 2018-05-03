package main

import (
	"testing"
)

/*
  comment
*/
func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()
	if expect != actual {
		t.Errorf("%s != %s", actual, expect)
	}
}
