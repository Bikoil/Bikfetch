package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
)

func main() {
	var uts syscall.Utsname
	syscall.Uname(&uts)

	osrelease, _ := os.Open("/etc/os-release")
	defer osrelease.Close()

	var osname string
	scanner := bufio.NewScanner(osrelease)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "PRETTY_NAME=") {
			osname = strings.Trim(strings.TrimPrefix(scanner.Text(), "PRETTY_NAME="), "\"")
			break
		} else if osname == "" {
	            osname = runtime.GOOS
	          }

	} 
	user := os.Getenv("USER")
	kernel := charsToString(uts.Release[:])
	host, _ := os.Hostname()
	display := os.Getenv("XDG_CURRENT_DESKTOP")

	fmt.Println("╭──────Hai There!─────╮",
		"\nOS —", osname,
		"\nKernel —", kernel,
		"\nUser —", user,
	        "\nHost Machine —", host,
		"\nWM / DE —", display,
		"\n╰─────────────────────╯")
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

