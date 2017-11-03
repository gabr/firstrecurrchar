package firstrecurrchar

import "testing"

func Test(t *testing.T) {
	var testCases = []struct {
		s    string
		want rune
	}{
		{"ABCA", 'A'},
		{"BCABA", 'B'},
		{"ABC", '\x00'},
	}

	for _, tc := range testCases {
		got := FindChar(tc.s)
		if got != tc.want {
			t.Errorf("FindChar(%q) == %q, want %q", tc.s, got, tc.want)
		}
	}
}
