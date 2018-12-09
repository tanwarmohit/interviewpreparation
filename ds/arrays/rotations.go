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

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func reverseRotate(a []int, d int) []int {
	n := len(a)
	d = d % n
	reverse(a[:d])
	reverse(a[d:])
	reverse(a)
	return a
}

func findPivot(a []int) int {
	low := 0
	high := len(a) - 1
	mid := 0

	for high > low {
		mid = low + (high-low)/2
		if low < mid && a[mid-1] > a[mid] {
			return mid - 1
		} else if mid < high && a[mid] > a[mid+1] {
			return mid
		} else if a[low] > a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearch(a []int, low, high, key int) int {
	mid := 0
	for high >= low {
		mid = low + (high-low)/2
		if key == a[mid] {
			return mid
		} else if key > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func findInRotatedSortedPivot(a []int, key int) int {
	n := len(a)
	if pivot := findPivot(a); pivot == -1 {
		return binarySearch(a, 0, n-1, key)
	} else if key == a[pivot] {
		return pivot
	} else if key >= a[0] {
		return binarySearch(a, 0, pivot-1, key)
	} else {
		return binarySearch(a, pivot+1, n-1, key)
	}
}

func findInRotatedSorted(a []int, key int) int {
	low, high, mid := 0, len(a)-1, 0
	for high > low {
		mid = low + (high-low)/2
		if key == a[mid] {
			return mid
		} else if a[low] <= a[mid] {
			if key >= a[low] && key <= a[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if key >= a[mid] && key <= a[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	if high == low && key == a[low] {
		return high
	}
	return -1
}

func findPairSumInSortedArray(a []int, sum int) (int, int) {
	low, high := 0, len(a)-1
	currSum := 0
	for high > low {
		currSum = a[low] + a[high]
		if currSum == sum {
			return a[low], a[high]
		} else if currSum > sum {
			high--
		} else {
			low++
		}
	}

	return -1, -1
}

func findPairSumInSortedRotatedArray(a []int, sum int) (int, int) {
	low, high := 0, 0
	if pivot := findPivot(a); pivot == -1 {
		low, high = 0, len(a)-1
	} else {
		low, high = pivot+1, pivot
	}
	n := len(a)
	currSum := 0
	for ((n - (low - high)) % n) != 0 {
		currSum = a[low] + a[high]
		if currSum == sum {
			return a[low], a[high]
		} else if currSum > sum {
			high = (n + (high - 1)) % n
		} else {
			low = (n + (low + 1)) % n
		}
	}

	return -1, -1
}

func insertionSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func maxSumRotationBrute(a []int) int {
	n := len(a)
	sum, maxSum := 0, 0
	for i := 0; i < n; i++ {
		a = juggling(a, i)
		sum = 0
		for j, v := range a {
			sum += j * v
		}
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}

func maxSumRotation(a []int) int {
	var sum, maxSum, arrSum int

	for i, v := range a {
		arrSum += v
		sum += i * v
	}

	maxSum = sum
	n := len(a)
	for i := 1; i < n; i++ {
		sum += n*a[i-1] - arrSum
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}

func findMinimumInRotatedSorted(a []int) int {
	low, high, mid := 0, len(a)-1, 0

	for high > low {
		mid = low + (high-low)/2
		if low < mid && a[mid-1] > a[mid] {
			return a[mid]
		} else if high > mid && a[mid] > a[mid+1] {
			return a[mid+1]
		} else if a[high] > a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return a[low]
}
