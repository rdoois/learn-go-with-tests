package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurtwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://blog.gypsydave5.com",
		"waat://furhurtwe.geds",
	}

	want := map[string]bool{
		"https://google.com":          true,
		"https://blog.gypsydave5.com": true,
		"waat://furhurtwe.geds":       false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
