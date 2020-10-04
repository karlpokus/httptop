package main

import (
  "os"
  "httptop"
)

func main() {
  httptop.Start(os.Stdin, os.Stdout)
}
