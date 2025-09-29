// go
package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cases := []struct {
		key string
		val []byte
	}{
		{"https://example.com", []byte("testdata")},
		{"https://example.com/path", []byte("moretestdata")},
	}

	for i, cse := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			cache.Add(cse.key, cse.val)
			got, ok := cache.Get(cse.key)
			if !ok {
				t.Fatalf("expected to find key")
			}
			if string(got) != string(cse.val) {
				t.Fatalf("unexpected value")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	interval := 5 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("k", []byte("v"))

	if _, ok := cache.Get("k"); !ok {
		t.Fatalf("expected key present")
	}

	time.Sleep(interval + 5*time.Millisecond)

	if _, ok := cache.Get("k"); ok {
		t.Fatalf("expected key reaped")
	}
}
