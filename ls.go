package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "."
	showAll := false
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			for _, c := range arg[1:] {
				if c == 'a' {
					showAll = true
				} else {
					fmt.Fprintf(os.Stderr, "usage: %s [-a] [path]\n", os.Args[0])
					os.Exit(1)
				}
			}
		} else {
			path = arg
		}
	}

	if showAll {
		fmt.Println(".")
		fmt.Println("..")
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, entry := range entries {
		if !showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		fmt.Println(entry.Name())
	}
}
