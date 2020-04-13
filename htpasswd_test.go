package htpasswd

import (
	"bytes"
	"reflect"
	"testing"
)

var source = []byte(`
johndoe:7fa37884
anonymous
# comment
janedoe:5281bb2c
`)

// TestNewReader ...
func TestNewReader(t *testing.T) {
	var want = [][]string{
		{"johndoe", "7fa37884"},
		{"janedoe", "5281bb2c"},
	}
	r := NewReader(bytes.NewReader(source))
	got, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v\nwant: %v", got, want)
	}
}
