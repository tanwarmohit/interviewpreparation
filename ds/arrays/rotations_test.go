package rotations

import "testing"

func deepIntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestGcd(t *testing.T) {
	a, b := 24, 44
	if output := gcd(a, b); output != 4 {
		t.Fatalf("wrong gcd(%d, %d) %v\n", a, b, output)
	}

	a, b = 44, 24
	if output := gcd(a, b); output != 4 {
		t.Fatalf("wrong gcd(%d, %d) %v\n", a, b, output)
	}

	a, b = 2, 9
	if output := gcd(a, b); output != 1 {
		t.Fatalf("wrong gcd(%d, %d) %v\n", a, b, output)
	}

	a, b = 2, 8
	if output := gcd(a, b); output != 2 {
		t.Fatalf("wrong gcd(%d, %d) %v\n", a, b, output)
	}
}

func TestJuggling(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	output := juggling(input, 2)
	if !deepIntArrayEquals(output, []int{3, 4, 5, 6, 7, 8, 1, 2}) {
		t.Fatalf("juggling failed\n")
	}

	output = juggling(input, 3)
	if !deepIntArrayEquals(output, []int{6, 7, 8, 1, 2, 3, 4, 5}) {
		t.Fatalf("juggling failed\n")
	}

	output = juggling(input[:2], 3)
	if !deepIntArrayEquals(output, []int{7, 6}) {
		t.Fatalf("juggling failed\n")
	}

}
