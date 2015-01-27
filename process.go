package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type factory interface {
	make(line string) task
}

type task interface {
	process()
	print()
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.make(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.print()
	}
}

// implementation of custom task

type testFactory struct {
}

type testTask struct {
	line   string
	err    error
	result int
}

func (t *testTask) process() {
	i, err := strconv.Atoi(t.line)
	if err != nil {
		t.err = err
		return
	}
	t.result = i * 10
}

func (t *testTask) print() {
	if t.err != nil {
		fmt.Println("Error", t.err)
		return
	}
	fmt.Println(t.result)
}

func (t *testFactory) make(line string) task {
	return &testTask{line: line}
}

func main() {
	run(&testFactory{})
}
