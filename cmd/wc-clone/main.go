package main

import (
	"flag"
	"fmt"
	"os"
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

	flag.Parse()

	files := flag.Args() // file names

	if len(files) == 0 {
		fmt.Println("File name is required!")
		os.Exit(1)
	}

	// output := internal.stats

	fmt.Println(*lines, *bytes, *chars)
}
