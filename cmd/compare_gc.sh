#!/bin/bash

echo "====================================================="
echo "        SCENARIO 1: The Ideal Case (Dense Heap)      "
echo "====================================================="

echo "Running Classic GC..."
GOEXPERIMENT=nogreenteagc go test -bench=^BenchmarkGC$ -count=10 > Classic_GC.txt

echo "Running Green Tea GC..."
GOEXPERIMENT=greenteagc go test -bench=^BenchmarkGC$ -count=10 > GreenTea_GC.txt

echo "====================================================="
echo "    SCENARIO 2: The Worst Case (Fragmented Heap)     "
echo "====================================================="

echo "Running Classic GC..."
GOEXPERIMENT=nogreenteagc go test -bench=^BenchmarkFragmentedHeap$ -count=10 > Classic_Frag.txt

echo "Running Green Tea GC..."
GOEXPERIMENT=greenteagc go test -bench=^BenchmarkFragmentedHeap$ -count=10 > GreenTea_Frag.txt


echo "====================================================="
echo "               STATISTICAL ANALYSIS                  "
echo "====================================================="

echo ""
echo "--- SCENARIO 1 (Dense Heap) ---"
~/go/bin/benchstat Classic_GC.txt GreenTea_GC.txt

echo ""
echo "--- SCENARIO 2 (Fragmented Heap) ---"
~/go/bin/benchstat Classic_Frag.txt GreenTea_Frag.txt