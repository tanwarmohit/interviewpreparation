package rotations

import (
	"testing"
)

func deepIntArrayEquals(a, b []int) bool {
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

func arrayCopy(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
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

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	reverse(input)
	if !deepIntArrayEquals(input, []int{4, 3, 2, 1}) {
		t.Fatalf("reverse failed\n")
	}

	input = []int{1, 2, 3}
	reverse(input)
	if !deepIntArrayEquals(input, []int{3, 2, 1}) {
		t.Fatalf("reverse failed\n")
	}
}

func TestReverseRotate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	output := reverseRotate(input, 2)
	if !deepIntArrayEquals(output, []int{3, 4, 5, 6, 7, 8, 1, 2}) {
		t.Fatalf("reverse rotate failed\n")
	}

	output = reverseRotate(input, 3)
	if !deepIntArrayEquals(output, []int{6, 7, 8, 1, 2, 3, 4, 5}) {
		t.Fatalf("reverse rotate failed\n")
	}

	output = reverseRotate(input[:2], 3)
	if !deepIntArrayEquals(output, []int{7, 6}) {
		t.Fatalf("reverse rotate failed\n")
	}
}

func TestPivot(t *testing.T) {
	if pivot := findPivot([]int{1, 2, 3, 4, 5, 6, 7}); pivot != -1 {
		t.Fail()
	}

	if pivot := findPivot([]int{7, 8, 1, 2}); pivot != 1 {
		t.Fail()
	}

	if pivot := findPivot([]int{7, 8, 1, 2, 3, 4, 5}); pivot != 1 {
		t.Fail()
	}

	if pivot := findPivot([]int{3, 4, 5, 6, 7, 8, 1, 2}); pivot != 5 {
		t.Fail()
	}

	if pivot := findPivot([]int{8, 1, 2}); pivot != 0 {
		t.Fail()
	}

	if pivot := findPivot([]int{8}); pivot != -1 {
		t.Fail()
	}

	if pivot := findPivot([]int{7, 8}); pivot != -1 {
		t.Fail()
	}

	if pivot := findPivot([]int{8, 7}); pivot != 0 {
		t.Fail()
	}
}

func TestBinarySearch(t *testing.T) {
	if out := binarySearch([]int{3, 5, 7, 8, 1, 2}, 0, 3, 98); out != -1 {
		t.Fail()
	}

	if out := binarySearch([]int{3, 5, 7, 8, 1, 2}, 4, 5, 2); out != 5 {
		t.Fail()
	}

	if out := binarySearch([]int{3, 5, 7, 8, 1, 2}, 0, 1, 5); out != 1 {
		t.Fail()
	}

	if out := binarySearch([]int{3, 5, 7, 8, 1, 2}, 1, 3, 17); out != -1 {
		t.Fail()
	}
}

func TestFindInRotatedSortedPivot(t *testing.T) {
	type input struct {
		data          []int
		key, expected int
	}
	inputData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	inputs := []input{
		{juggling(arrayCopy(inputData), 3), 1, 6},
		{juggling(arrayCopy(inputData), 0), 1, 0},
		{juggling(arrayCopy(inputData), 7), 2, 3},
		{juggling(arrayCopy(inputData[:3]), 1), 3, 1},
		{juggling(arrayCopy(inputData), 5), 6, 0},
	}

	for _, data := range inputs {
		if actual := findInRotatedSortedPivot(data.data, data.key); actual != data.expected {
			t.Fail()
		}
	}
}

func TestFindInRotatedSorted(t *testing.T) {
	type input struct {
		data          []int
		key, expected int
	}
	inputData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	inputs := []input{
		{juggling(arrayCopy(inputData), 3), 1, 6},
		{juggling(arrayCopy(inputData), 0), 1, 0},
		{juggling(arrayCopy(inputData), 7), 2, 3},
		{juggling(arrayCopy(inputData[:3]), 1), 3, 1},
		{juggling(arrayCopy(inputData), 5), 6, 0},
		{juggling(arrayCopy(inputData), 7), 4, 5},
		{juggling(arrayCopy(inputData[:1]), 7), 1, 0},
		{juggling(arrayCopy(inputData[:1]), 7), 11, -1},
	}

	for _, data := range inputs {
		if actual := findInRotatedSorted(data.data, data.key); actual != data.expected {
			t.Log(actual, data.expected)
			t.Fail()
		}
	}
}

func TestInsertionSort(t *testing.T) {

	type sortInput struct {
		input  []int
		output []int
	}

	inputs := []sortInput{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{4, 2, 5, 8, 98, 45, 12}, []int{2, 4, 5, 8, 12, 45, 98}},
	}

	for _, data := range inputs {
		if insertionSort(data.input); !deepIntArrayEquals(data.input, data.output) {
			t.Fail()
		}
	}

}
func TestFindPairSumInSortedArray(t *testing.T) {
	type pairSumInput struct {
		input                         []int
		sum                           int
		expectedFirst, expectedSecond int
	}

	pairSumInputs := []pairSumInput{
		{[]int{2, 5, 4, 12, 1, 14}, 16, 2, 14},
		{[]int{2, 5, 4, 12, 1, 14}, 9, 4, 5},
		{[]int{2, 5, 4, 12, 1, 14}, 29, -1, -1},
	}

	for _, input := range pairSumInputs {
		insertionSort(input.input)
		if actualFirst, actualSecond := findPairSumInSortedArray(input.input, input.sum); actualFirst != input.expectedFirst || actualSecond != input.expectedSecond {
			t.Logf("af: %d, as: %d\n", actualFirst, actualSecond)
			t.Fail()
		}
	}
}

func TestFindPairSumInSortedRotatedArray(t *testing.T) {
	type pairSumInput struct {
		input                         []int
		rotationCount                 int
		sum                           int
		expectedFirst, expectedSecond int
	}

	pairSumInputs := []pairSumInput{
		{[]int{2, 5, 4, 12, 1, 14}, 2, 16, 2, 14},
		{[]int{2, 5, 4, 12, 1, 14}, 2, 9, 4, 5},
		{[]int{2, 5, 4, 12, 1, 14}, 0, 29, -1, -1},
	}

	for _, input := range pairSumInputs {
		insertionSort(input.input)
		input.input = juggling(input.input, input.rotationCount)
		if actualFirst, actualSecond := findPairSumInSortedRotatedArray(input.input, input.sum); actualFirst != input.expectedFirst || actualSecond != input.expectedSecond {
			t.Logf("af: %d, as: %d, ef: %d, es: %d\n", actualFirst, actualSecond, input.expectedFirst, input.expectedSecond)
			t.Fail()
		}
	}
}

func TestMaxSumRotation(t *testing.T) {
	inputs := [][]int{
		{2, 5, 6, 3, 5, 12, 78},
		{5, 4, 3, 2, 1},
	}

	for _, input := range inputs {
		if maxSumRotation(input) != maxSumRotationBrute(input) {
			t.Fail()
		}
	}
}
