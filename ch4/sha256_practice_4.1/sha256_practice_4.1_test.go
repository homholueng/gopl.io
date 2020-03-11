package shadiff

import (
	"testing"
)

func TestSha256DiffBits(t *testing.T) {
	a1 := [32]byte{}
	a1[0] = 0

	b1 := [32]byte{}
	b1[0] = 6

	a2 := [32]byte{}
	a2[0] = 1
	a2[1] = 2
	a2[2] = 3

	b2 := [32]byte{}
	b2[0] = 4
	b2[1] = 5
	b2[2] = 6

	tests := []struct {
		a, b [32]byte
		want uint
	}{
		{a1, b1, 2},
		{a2, b2, 7},
	}

	for _, test := range tests {
		got := Sha256DiffBits(test.a, test.b)
		if got != test.want {
			t.Errorf("Sha256DiffBits(%v, %v) got %d, want %d",
				test.a, test.b, got, test.want)
		}
	}
}
