package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "watt://down.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://xvidesos.com",
		"watt://down.com",
	}

	want := map[string]bool{
		"https://google.com":   true,
		"https://xvidesos.com": true,
		"watt://down.com":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v; got %v\n", want, got)
	}
}
