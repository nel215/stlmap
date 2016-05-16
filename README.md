[![Build Status](https://travis-ci.org/nel215/stlmap.svg?branch=master)](https://travis-ci.org/nel215/stlmap)

# About

stlmap is a concurrent map using striped locked approach

# Example

see [here](./example_test.go)

# Performance

```
BenchmarkStlmapConcurrent-4             20000000               257 ns/op
BenchmarkDefaultMapConcurrent-4         20000000               315 ns/op
BenchmarkStlmapConcurrentSet-4          20000000               211 ns/op
BenchmarkDefaultMapConcurrentSet-4      20000000               223 ns/op
BenchmarkStlmapConcurrentGet-4          30000000               142 ns/op
BenchmarkDefaultMapConcurrentGet-4      20000000               161 ns/op
BenchmarkStlmapConcurrentDelete-4       30000000               154 ns/op
BenchmarkDefaultMapConcurrentDelete-4   30000000               167 ns/op
```
