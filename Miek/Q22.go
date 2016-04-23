/* Mimic the "cat" program */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
CAVEAT: flag lib only work if flag(s) ahead of arg(s)
otherwise, it will think flag(s) is arg(s)
TODO: Will try other lib to make this "cli" works with flag(s) in any position
*/
var numbered = flag.Bool("n", false, "numbered for each line")

func main() {
	flag.Parse()
	fmt.Printf("Value of 'n' flag = %v\n", *numbered)
	if flag.NArg() == 0 {
		fmt.Println("Missing filename. Pls use: ./Q22.go <path/to/file> [--n=<true|false>]")
		return
	} else {
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i)) // f is *os.File
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
				continue // ?? why need this
			}
			cat(bufio.NewReader(f))
		}

	}

}

func cat(r *bufio.Reader) {
	i := 1
	// Improved to handle file with non '\n' ending line(s)
	for line, isPrefix, err := r.ReadLine(); err == nil; line, isPrefix, err = r.ReadLine() {
		if *numbered {
			fmt.Fprintf(os.Stdout, "%4d", i)
			i++
		}
		fmt.Fprintf(os.Stdout, string(line))
		// Just try "Fprintf"; in practical we should use fmt.Printf()

		if isPrefix { // mean buffer is not enough to read this line
			for {
				line, isPrefix, _ = r.ReadLine()
				fmt.Fprintf(os.Stdout, string(line))
				if !isPrefix {
					break
				}
			}
		}
		fmt.Fprintf(os.Stdout, "\n")
	}

}
