# Go GC Comparison: Classic vs Green Tea

Companion repo for a Medium article comparing the **classic Go GC** with the **Green Tea GC** (Go 1.26).

## Run the benchmarks

**Requirements:** Go 1.26, [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) (`go install golang.org/x/perf/cmd/benchstat@latest`). The script uses `~/go/bin/benchstat`—adjust `compare_gc.sh` if yours is elsewhere.

```bash
cd cmd
./compare_gc.sh
```

Trace + manual GC (optional): `go run ./cmd` then `go tool trace trace.out`.

---

## Example results (Apple M4, arm64, darwin)

### Scenario 1 — Dense heap

| Metric             | Classic      | Green Tea    | vs base   |
|--------------------|--------------|--------------|-----------|
| **sec/op**         | 31.01m ± 5%  | 23.87m ± 21% | -23.02%   |
| **MB_Given_Back**  | 5.699 ± 7%   | 9.270 ± 3%   | +62.65%   |
| **MB_Peak_Memory** | 520.9 ± 0%   | 533.6 ± 0%   | +2.44%    |
| **B/op**           | 1.032Ki ± 1% | 1.000Ki ± 0% | -3.08%    |

### Scenario 2 — Fragmented heap

| Metric             | Classic       | Green Tea     | vs base   |
|--------------------|---------------|---------------|-----------|
| **sec/op**         | 3.864m ± 61%  | 1.528m ± 37%  | -60.47%   |
| **MB_Peak_Memory** | 91.97 ± 90%   | 12.82 ± 145%  | -86.05%   |
| **B/op**           | 1.000Ki ± 0%  | 1.000Ki ± 0%  | ~         |

---

[LICENSE](LICENSE)
