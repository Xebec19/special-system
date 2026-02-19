package main

import (
	"flag"
	"fmt"
)

/*
*
implement -c, -m and -l first
*/
func main() {
	// Raw args: os.Args[0] is program name

	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("c", false, "count bytes")
	chars := flag.Bool("m", false, "count characters")

	fmt.Println(lines, bytes, chars)
}
