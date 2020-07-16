package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open("myfile.data")
	defer f.Close()
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened sucessfully")

	//	Creating dnsError Struct with error

	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)

	//	Using Gob function of the filepath package to return the names
	//	of the all the files that matches a pattern
	files, err := filepath.Glob("[")
	if err != nil {
		//	ErrBadPattern is returned by the Glob function when the pattern is malformed.
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)

}
