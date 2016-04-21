package main

// No such thing as "from foo import bar"
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Note how the type comes after arg name in argument list
func countWords(source io.Reader) (map[string]int, error) {
	counts := make(map[string]int) // What's this?
	scanner := bufio.NewScanner(bufio.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++ // has ++ but it's not crazy like C
	}
	// returning errors is a common pattern
	return counts, scanner.Err()
}

func main() {
	counts, err := countWords(os.Stdin) // type inference!
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for word, count := range counts { // "for" does all types of loops
		if count > 1 {
			// print anything, just like Python
			fmt.Println(word, count)
		}
	}
}
