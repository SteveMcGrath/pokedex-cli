package pokecache

import (
	"bytes"
	"testing"
)

func TestCacheAdd(t *testing.T) {
	c := NewCache()

	if _, ok := c.Get("test"); ok {
		t.Errorf("new cache, should not have an entry")
	}

	tobj := []byte{'H', 'e', 'l', 'l', 'o'}
	c.Add("test", tobj)

	if obj, ok := c.Get("test"); !ok {
		t.Errorf("test added to cache, however did not return")
	} else if !bytes.Equal(obj, tobj) {
		t.Errorf("test object doesn't match the retreived cache")
	}
}

func TestCacheGetJSON(t *testing.T) {
	c := NewCache()
	if _, ok := c.Get("test"); ok {
		t.Errorf("new cache, should not have an entry")
	}

	data := struct {
		Args    map[string]any `json:"args"`
		Headers map[string]any `json:"headers"`
		Origin  string         `json:"origin"`
		Url     string         `json:"url"`
	}{}
	if err := c.GetJSON("https://httpbin.org/get", &data); err != nil {
		t.Error(err)
	}
}
