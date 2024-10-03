package main

import (
  "fmt"
  "os"
  "runtime"
)

func main() {
  user := os.Getenv("USER")
  osname := runtime.GOOS
  fullosname :=


  fmt.Println("OS:", osname, 
  "\nUser:", user)
}
