package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rahji/wordsearch"
	"github.com/rahji/wordsearch-cli/internal/wordlist"
	"github.com/spf13/pflag"
)

func main() {

	// define flags
	var (
		inputFile string
		size      int
		help      bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.IntVarP(&size, "size", "s", 16, "grid size (default: 16)")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	words, err := wordlist.GetWords(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ws := wordsearch.NewWordSearch(size, nil, true)
	unplaced := ws.CreatePuzzle(words)

	// make a map of the unplaced words for later lookup
	unplacedMap := make(map[string]struct{})
	for _, word := range unplaced {
		unplacedMap[word] = struct{}{}
	}

	bold := color.New(color.FgBlue)

	grid := ws.ReturnGrid(wordsearch.GridRaw)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 97 { // lowercase
				fmt.Printf("%c ", grid[i][j]-32)
			} else { // uppercase
				bold.Printf("%c ", grid[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println()
	for _, w := range words {
		if _, exists := unplacedMap[w]; !exists {
			fmt.Println(w)
		}
	}

	if unplaced != nil {
		fmt.Printf("\nFYI these %d words could not be placed: %v\n", len(unplaced), unplaced)
	}

}
