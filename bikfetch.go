package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
	"errors"
)

func main() {
	var uts syscall.Utsname
	syscall.Uname(&uts)
    // Get os release path variable
    var releasepath string = "/etc/os-release" 
	osrelease, _ := os.Open("/etc/os-release")
	defer osrelease.Close()

	var osname string
	scanner := bufio.NewScanner(osrelease)
    // OSNAME will be the normal runtime os name if /etc/os-release doesnt work
    osname = runtime.GOOS
	for scanner.Scan() {
		if checkOsRelease(releasepath) && strings.HasPrefix(scanner.Text(), "PRETTY_NAME=") {
			osname = strings.Trim(strings.TrimPrefix(scanner.Text(), "PRETTY_NAME="), "\"")
			break
		} 

	} 
	user := os.Getenv("USER") // User
	kernel := charsToString(uts.Release[:]) // Kernel
	host, _ := os.Hostname() // Hostname
	display := os.Getenv("XDG_CURRENT_DESKTOP") // Display (WM/DE)
	if os.Getenv("XDG_DESKTOP_SESSION") == "" {
          display = os.Getenv("DESKTOP_SESSION")
	}
	if os.Getenv("DESKTOP_SESSION") == "" {
            display = os.Getenv("GDMSESSION") 
	}

	// Print it all out
	fmt.Println("╭──────Hai There!─────╮",
		"\nOS —", osname,
		"\nKernel —", kernel,
		"\nUser —", user,
	        "\nHost —", host,
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
func checkOsRelease(releasepath string) bool {
	_, error := os.Stat(releasepath)
	return !errors.Is(error, os.ErrNotExist)
}
