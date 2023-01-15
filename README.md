# Password
Password is a password validation package written in Go used for validating passwords in the backend

## Getting started
```shell
go get github.com/marcoshuck/password
```

## Usage
```go
err := password.Validate("%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")

if err != nil {
  log.Println("Failed to validate password:", err)
}
```

Play with this package in this [Playground](https://go.dev/play/p/zAhLEuKnmZ1) link.


## Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/marcoshuck/password/v1
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkValidatePassword
BenchmarkValidatePassword-12             2634285               457.5 ns/op
PASS
```
