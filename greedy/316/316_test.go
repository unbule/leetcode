package greedy

import (
	"testing"
)

func Test_316(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect string
		real   string
	}{
		{"case1", "bcabc", "abc", ""},
		{"case2", "cbacdcbc", "acdb", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.real = removeDuplicateLetters(tt.in)
			if tt.expect != tt.real {
				t.Fatal("expect:", tt.expect, "real:", tt.real)
			}
		})
	}
}
