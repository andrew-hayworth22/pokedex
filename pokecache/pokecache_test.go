package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestPokecache(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	cachedValues := []struct {
		Name string
		Val  []byte
	}{
		{
			"https://google.com",
			[]byte("API Response"),
		},
		{
			"https://google.com/path",
			[]byte("API Response 2"),
		},
	}

	for _, cachedValue := range cachedValues {
		cache.Add(cachedValue.Name, cachedValue.Val)
		val, ok := cache.Get(cachedValue.Name)
		if !ok || !bytes.Equal(val, cachedValue.Val) {
			t.Errorf("%s key not found in cache", cachedValue.Name)
		}
	}

	noval, ok := cache.Get("https://doesntexist.com")
	if noval != nil || ok {
		t.Errorf("key not supposed to be found(val=%v ok=%t)", noval, ok)
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + (5 * time.Millisecond)

	cache := NewCache(interval)
	cachedValues := []struct {
		Name string
		Val  []byte
	}{
		{
			"https://google.com",
			[]byte("API Response"),
		},
		{
			"https://google.com/path",
			[]byte("API Response 2"),
		},
	}

	for _, cachedValue := range cachedValues {
		cache.Add(cachedValue.Name, cachedValue.Val)
	}

	time.Sleep(waitTime)

	for _, val := range cachedValues {
		cached, ok := cache.Get(val.Name)
		if cached != nil || ok {
			t.Errorf("cached value not reaped (val=%v ok=%t)", cached, ok)
		}
	}
}
