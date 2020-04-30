package check_odd

import "testing"

func TestCheckOdd(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  bool
	}{
		{"Expect true", 0, true},
		{"Expect false", 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testCheckOdd(t, test.in, test.out)
		})
	}
}

func testCheckOdd(t *testing.T, in int, expected bool) {
	r := CheckOdd(in)
	if r != expected {
		t.Errorf("CheckOdd(%d) = %t, want %t", in, r, expected)
	}
}
