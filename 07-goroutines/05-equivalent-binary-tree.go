package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

////////////////////////////////////////
func Walk(t *Tree, ch chan int) {
    WalkRecursive(t, ch)
    close(ch)
}

func WalkRecursive(t *Tree, ch chan int) {
    if t != nil {
        WalkRecursive(t.Left, ch)
        ch <- t.Value
        WalkRecursive(t.Right, ch)
    }
}

func Same(t1, t2 *Tree) bool {
    ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for {
        n1, ok1 := <- ch1
        n2, ok2 := <- ch2
        if ok1 != ok2 || n1 != n2 {
        	return false
        }
        if !ok1 {
        	break;
        }
    }
    return true
    
    
}

func main() {
    ch := make(chan int)
    t := New(1)
    fmt.Println(t.String())
    go Walk(t, ch)
    fmt.Println()
    fmt.Println(Same(New(1), New(1)))
    fmt.Println(Same(New(1), New(2)))
}