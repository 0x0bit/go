package main

import (
	"testing"
)

func TestSubStr(t *testing.T) {
	tests :=[] struct{
		s string
		ans int
	}{
		{"aaaabcabc", 3},
		{"bbbbbbbbbbbb", 1},
		{"一二三二一", 3},
		{"我爱大中国", 5},
		{"灰化肥会挥发", 0},
	}

	for _, tt := range tests {
		actual := lengthOfSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}
