package rotations

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// Juggling ... Perfoms left rotation of array by d steps
func juggling(a []int, d int) []int {
	if d == 0 {
		return a
	}

	n := len(a)
	d = d % n
	m := gcd(n, d)
	var k, next, prev int
	for i := 0; i < m; i++ {
		k = i
		prev = a[i]
		for {
			next = (k + d) % n
			if next == i {
				break
			}
			a[k] = a[next]
			k = next
		}
		a[k] = prev
	}
	return a
}
