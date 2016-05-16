[![Build Status](https://travis-ci.org/nel215/stlmap.svg?branch=master)](https://travis-ci.org/nel215/stlmap)

# About

stlmap is a concurrent map using striped locked approach

# Example

see [here](./example_test.go)

# Performance

```
BenchmarkStlmapConcurrent-4     20000000               260 ns/op
BenchmarkDefaultMapConcurrent-4 20000000               310 ns/op
```
