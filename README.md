### gRPC+Protobuf or JSON+HTTP or JSON+(FastHttp Based)?

This repository contains 3 equal APIs: gRPC using Protobuf, JSON over go HTTP and JSON over FastHttp Based. The goal is to run benchmarks for 3 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 3 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.18

### Run tests

Run benchmarks:
```
GO111MODULE=on go test -bench=. -benchmem -cpu=1
```

We are using only 1 cpu to avoid too much routine overhead in so fast methods.

### Results

```
goos: darwin
goarch: amd64
pkg: benchmark-grpc-protobuf-vs-http-json
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkFiberJSON         12621             96835 ns/op           11058 B/op        115 allocs/op
BenchmarkGRPCProtobuf      11956            102277 ns/op           13907 B/op        209 allocs/op
BenchmarkHTTPJSON          11223            103864 ns/op           14860 B/op        143 allocs/op
PASS
ok      benchmark-grpc-protobuf-vs-http-json    9.725s
```

They are almost the same. Bigger is the payload, faster grpc will be respect to others.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```
GO111MODULE=on go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu -cpu=1
GO111MODULE=on go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu -cpu=1
GO111MODULE=on go test -bench=BenchmarkFiberJSON -cpuprofile=fiberjson.cpu -cpu=1
```

Check CPU usage per approach using:

```
go tool pprof grpcprotobuf.cpu
go tool pprof httpjson.cpu
```


### gRPC definition

 - Install [Go](https://golang.org/dl/)
 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)

```
cd grpc-protobuf
make
```