package day3

import (
	"testing"

	"github.com/wincus/adventofcode2021/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"00100", //4
				"11110", //30
				"10110", //22
				"10111", //23
				"10101", //21
				"01111", //15
				"00111", //7
				"11100", //28
				"10000", //16
				"11001", //25
				"00010", //2
				"01010", //10
			},
			p:    common.Part1,
			want: 198,
		},
		{
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			p:    common.Part2,
			want: 230,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
