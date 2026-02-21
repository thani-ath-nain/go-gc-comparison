package tests

import (
	"runtime"
	"testing"
)

// Node is our test object
type Node struct {
	Children [4]*Node
	Data     [8]int64
}

func createTree(depth int) *Node {
	if depth <= 0 {
		return &Node{}
	}
	n := &Node{}
	for i := range n.Children {
		n.Children[i] = createTree(depth - 1)
	}
	return n
}

func BenchmarkGC(b *testing.B) {
	// Setup: Create the heap pressure once
	root := createTree(11)
	runtime.GC() // Clean up setup noise
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.GC()
	}

	runtime.KeepAlive(root)
}
