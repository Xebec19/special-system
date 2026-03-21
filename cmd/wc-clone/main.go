package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Xebec19/special-system/internal"
	"github.com/Xebec19/special-system/internal/version"
)

/*
*
implement -c, -m and -l first
*/
func main() {
	// Raw args: os.Args[0] is program name

	showVersion := flag.Bool("version", false, "print version")
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("c", false, "count bytes")
	chars := flag.Bool("m", false, "count characters")
	words := flag.Bool("w", false, "count words")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\ncommit: %s\nbuilt: %s\n",
			version.Version,
			version.Commit,
			version.BuildDate,
		)
		return
	}

	files := flag.Args() // file names

	if len(files) == 0 {
		fmt.Println("File name is required!")
		os.Exit(1)
	}

	output := internal.Stats(*lines, *chars, *words, *bytes, files)

	fmt.Print(output)
}
