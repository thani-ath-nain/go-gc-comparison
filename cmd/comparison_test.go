package main

import (
	"runtime"
	"testing"
)

var sink []byte

func BenchmarkGC(b *testing.B) {
	root := CreateTree(11)
	runtime.GC()
	b.ReportAllocs()
	for b.Loop() {
		sink = make([]byte, 1024)
		runtime.GC()
	}
	b.StopTimer()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	b.ReportMetric(float64(m.HeapInuse)/1024/1024, "MB_Peak_Memory")
	b.ReportMetric(float64(m.HeapReleased)/1024/1024, "MB_Given_Back")
	runtime.KeepAlive(root)
}

func BenchmarkFragmentedHeap(b *testing.B) {
	var head *SmallNode
	var curr *SmallNode
	for range 50000 {
		// Allocate a batch of 170 objects (approx. one 8KB span)
		batch := make([]*SmallNode, 170)
		for j := range 170 {
			batch[j] = &SmallNode{}
		}
		// Keep exactly ONE object per 8KB page alive
		nodeToKeep := batch[85]
		if head == nil {
			head = nodeToKeep
			curr = head
		} else {
			curr.Next = nodeToKeep
			curr = curr.Next
		}
	}
	runtime.GC()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = make([]byte, 1024)
		runtime.GC()
	}
	b.StopTimer()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	b.ReportMetric(float64(m.HeapInuse)/1024/1024, "MB_Peak_Memory")
	runtime.KeepAlive(head)
}
