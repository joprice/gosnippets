package main

// example usage: echo -i "abc\ndef" | go run stdin.go

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println("got line", s.Text())
	}
	if s.Err() != nil {
		log.Fatalf("Error reading STDIN: %s", s.Err())
	}
}
