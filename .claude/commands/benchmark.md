Run Go benchmarks to measure logger performance.

Execute:

```bash
cd /Users/olegiv/Desktop/Projects/Go/go-logger

# Run benchmarks
go test -bench=. -benchmem -run=^$
```

Display:
- Benchmark results with iterations/second
- Memory allocations per operation
- Bytes allocated per operation

Note: If no benchmarks exist yet, suggest creating benchmark tests for:
- Logger creation (`BenchmarkNew`)
- Basic logging operations (`BenchmarkInfoLog`)
- Logging with fields (`BenchmarkWithFields`)
- Context methods (`BenchmarkWithField`, `BenchmarkWithError`)

Example benchmark format:
```go
func BenchmarkInfoLog(b *testing.B) {
    tmpDir := b.TempDir()
    cfg := Config{Level: "info", LogDir: tmpDir}
    log := New(cfg)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        log.Info().Msg("benchmark message")
    }
}
```

Benchmarks help ensure the logger maintains zero-allocation performance characteristics.
