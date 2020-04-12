package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return fmt.Sprintf("%v", dig(t, make([]int, 0)))
}

func dig(t *tree, values []int) []int {
	if t.left != nil {
		values = dig(t.left, values)
	}
	values = append(values, t.value)
	if t.right != nil {
		values = dig(t.right, values)
	}
	return values
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	t := &tree{value: 0}
	add(t, 1)
	add(t, -2)
	add(t, 3)
	fmt.Println(t)
}