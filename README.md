# ascii
# [![GoDoc](https://godoc.org/github.com/thanhpk/ascii?status.svg)](http://godoc.org/github.com/thanhpk/ascii)

Converts an unicode string to an ascii string

## Install
```
  go get -u github.com/thanhpk/ascii
```

## Usage
```go
out := ascii.Convert("Trời hôm nay đẹp quá") // Troi hom nay dep qua
```
Running example
```go
  package main
  import(
    "github.com/thanhpk/ascii"
    "fmt"
  )

  func main() {
    fmt.Println(ascii.Convert("Nam quốc sơn hà mam đế cư"))
    fmt.Println(ascii.Convert("Cuántos años tienes"))
    fmt.Println(ascii.Convert("AppleInc 是苹果公司."))
  }
  // Output:
  // Nam quoc son ha mam de cu
  // Cuantos anos tienes
  // AppleInc .

```

## License [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
MIT
