package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

// Would you like me to show you how to add runtime.MemStats
// to this script so you can also report on "Percentage Reduction in Peak Memory"?
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

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	fmt.Println("Starting GC Locality Benchmark...")
	var ms runtime.MemStats
	start := time.Now()
	root := createTree(11)
	fmt.Printf("Allocation took: %v\n", time.Since(start))
	for i := range 5 {
		gcStart := time.Now()
		runtime.GC()
		fmt.Printf("Manual GC %d took: %v\n", i+1, time.Since(gcStart))
	}
	runtime.ReadMemStats(&ms)
	fmt.Printf("Final HeapAlloc: %d MB\n", ms.HeapAlloc/1024/1024)
	fmt.Printf("Total GC Pause: %v\n", time.Duration(ms.PauseTotalNs))
	runtime.KeepAlive(root)
}
