package main

import (
  "fmt"
  "os"
  "runtime"
  "syscall"
)

func main() {

  var uts syscall.Utsname
  syscall.Uname(&uts)

//  file, _ := os.Open("/etc/os-release")
  user := os.Getenv("USER")
  osname := runtime.GOOS
  kernel:= charsToString(uts.Release[:])


  fmt.Println("OS:", osname,
  "\nKernel:", kernel, 
  "\nUser:", user)
}

// FUNctions... yaaay
  func charsToString(ca []int8) string {
	  s := make([]byte, len(ca))
	  for i, v := range ca {
		  if v == 0 {
			  break
		  }
		  s[i] = byte(v)
	  }
          return string(s)
  }
