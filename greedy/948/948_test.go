package greedy

import "testing"

func Test_948(t *testing.T) {
	tests := []struct {
		name   string
		tokens []int
		power  int
		expect int
		real   int
	}{
		{name: "case1", tokens: []int{100}, power: 50, expect: 0, real: 0},
		{name: "case2", tokens: []int{100, 200}, power: 150, expect: 1, real: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.real = bagOfTokensScore(test.tokens, test.power)
			if test.real != test.expect {
				t.Fatal("test.real:", test.real, "test.expect", test.expect)
			}
		})
	}
}
