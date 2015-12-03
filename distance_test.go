package distance

import "testing"

func TestLevenshteinDistanceIsValid(t *testing.T) {
	samples := []struct {
		first    string
		second   string
		expected int
	}{
		{"hello", "hello", 0},
		{"hello", "hallo", 1},
		{"s", "x", 1},
		{"Hello", "hello", 1},
		{"Sup", "Sup\n", 1},
		{"spartan", "sports", 3},
		{"axxxxx", "xxxxxb", 2},
	}
	// Test using all of the sample data.
	for _, sample := range samples {
		dist := Distance(sample.first, sample.second)
		if dist != sample.expected {
			t.Errorf("expected: %d, got %d", sample.expected, dist)
		}
	}
}
