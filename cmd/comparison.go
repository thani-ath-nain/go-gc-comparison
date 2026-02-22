package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

type Node struct {
	Children [4]*Node
	Data     [8]int64
}

type SmallNode struct {
	Next *SmallNode
	Data [5]int64
}

func CreateTree(depth int) *Node {
	if depth <= 0 {
		return &Node{}
	}
	n := &Node{}
	for i := range n.Children {
		n.Children[i] = CreateTree(depth - 1)
	}
	return n
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\n--- %s ---\n", label)
	fmt.Printf("HeapAlloc:    %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("HeapInuse:    %v MB\n", m.HeapInuse/1024/1024)
	fmt.Printf("HeapIdle:     %v MB\n", m.HeapIdle/1024/1024)
	fmt.Printf("HeapReleased: %v MB\n", m.HeapReleased/1024/1024)
	fmt.Printf("NumGC:        %v\n", m.NumGC)
}

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	fmt.Println("Starting GC Locality Benchmark...")
	start := time.Now()
	root := CreateTree(11)
	fmt.Printf("Allocation took: %v\n", time.Since(start))
	printMemStats("Post-Allocation")
	for i := range 5 {
		gcStart := time.Now()
		runtime.GC()
		fmt.Printf("Manual GC %d took: %v\n", i+1, time.Since(gcStart))
	}
	printMemStats("Final State")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\nTotal GC Pause: %v\n", time.Duration(m.PauseTotalNs))
	runtime.KeepAlive(root)
}
