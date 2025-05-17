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

	// OSNAME will be the normal runtime os name if /etc/os-release doesn't work
	osname := runtime.GOOS
	if checkOsRelease(releasepath) {
		osrelease, err := os.Open(releasepath)
		if err == nil {
			defer osrelease.Close()
			scanner := bufio.NewScanner(osrelease)
			for scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), "PRETTY_NAME=") {
					osname = strings.Trim(strings.TrimPrefix(scanner.Text(), "PRETTY_NAME="), "\"")
					break
				}
			}
		}
	}
	if osname == "" {
		osname = "N/A"
	}

	user := os.Getenv("USER") // User
	if user == "" {
		user = "N/A"
	}

	kernel := charsToString(uts.Release[:]) // Kernel
	if kernel == "" {
		kernel = "N/A"
	}

	host, err := os.Hostname() // Hostname
	if err != nil || host == "" {
		host = "N/A"
	}

	display := os.Getenv("XDG_CURRENT_DESKTOP") // Display (WM/DE)
	if display == "" {
		display = os.Getenv("XDG_DESKTOP_SESSION")
	}
	if display == "" {
		display = os.Getenv("GDMSESSION") 
	}
	if display == "" {
		display = "N/A"
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
