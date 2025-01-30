package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 15 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("example data"),
		},
		{
			key: "https://example2.com/somepath",
			val: []byte("more test data"),
		},
		{
			key: "https://example3.com/somepath/somemorethings",
			val: []byte("even more test data 234"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {

			cache := NewCache(interval)

			cache.Add(c.key, c.val)

			data, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("Expected to be found in cache, but it wasn't")
				return
			}

			if string(data) != string(c.val) {
				t.Errorf("Expected data to be %s, got %s", string(c.val), string(data))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	testKey := "https://exampletest.com"
	testData := []byte("test data")

	cache := NewCache(baseTime)
	cache.Add(testKey, testData)

	_, exists := cache.Get(testKey)
	if !exists {
		t.Errorf("Expected to be found in cache, but it wasn't")
		return
	}

	time.Sleep(waitTime)
	_, exists = cache.Get(testKey)
	if exists {
		t.Errorf("Expected to not be found in cache, but it was")
	}
}
