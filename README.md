# Password
Password is a password validation package written in Go used for validating passwords in the backend

## Getting started
```shell
go get github.com/marcoshuck/password/v1
```

## Usage
```go
err := password.Validate("%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")

if err != nil {
  log.Println("Failed to validate password:", err)
}
```
