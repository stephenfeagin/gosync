package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: gosync [src] [dst]")
		os.Exit(1)
	}
	src, dst := os.Args[1], os.Args[2]
	_, err := Copy(src, dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
