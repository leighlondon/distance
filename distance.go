// Package distance calculates the edit distance between strings.
package distance

func Distance(s1, s2 string) int {
	// Check whether the strings are equal.
	if s1 == s2 {
		return 0
	}
	// Check for empty strings.
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}
	// Check whether the last character is the same.
	cost := 1
	if s1[len(s1)-1] == s2[len(s2)-1] {
		cost = 0
	}
	// Calculate the edit distance recursively.
	// Take the minimum of three different distances.
	return min(min(
		Distance(s1, s2[0:len(s2)-1])+1,
		Distance(s1[0:len(s1)-1], s2)+1),
		Distance(s1[0:len(s1)-1], s2[0:len(s2)-1])+cost)
}

// min returns the smaller of two integers.
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
