package goshort

import "testing"

var encodeTests = []struct {
	id       int    // input
	expected string // expected result
}{
	{3141592, "vJST"},
	{123456789, "pgK8p"},
}

func TestEncode(t *testing.T) {
	for _, tt := range encodeTests {
		actual := Encode(tt.id)
		if actual != tt.expected {
			t.Error(
				"For", tt.id,
				"expected", tt.expected,
				"got", actual,
			)
		}
	}
}

func TestDecode(t *testing.T) {
	var id int
	id, _ = Decode("vJST")
	if id != 3141592 {
		t.Error("Expected 3141592, got ", id)
	}
}
