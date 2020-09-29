package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// replace fields with any type and modify comparison operators accordingly
type pair struct {
	x int
	y int
}

type pairs []pair

func (a pairs) Len() int           { return len(a) }
func (a pairs) Less(i, j int) bool { return a[i].x < a[j].x || a[i].y < a[j].y }
func (a pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// file
	// ----
	f, err := os.Open("a.in")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		// NOTE here
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// stdio
	// -----
	defer writer.Flush()

	// NOTE here
	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("%d\n", a+b)
}
