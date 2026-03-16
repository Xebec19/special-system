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
	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\ncommit: %s\nbuilt: %s\n",
			version.Version,
			version.Commit,
			version.BuildDate,
		)
		return
	}

	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("c", false, "count bytes")
	chars := flag.Bool("m", false, "count characters")

	flag.Parse()

	files := flag.Args() // file names

	if len(files) == 0 {
		fmt.Println("File name is required!")
		os.Exit(1)
	}

	output := internal.Stats(*lines, *chars, *bytes, files)

	fmt.Print(output)
}
