package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  filepath.Walk("C:\\gowork\\src\\github.com\\wagnergaldino\\learning\\anintroductiontoprogrammingingo", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
  })
}
