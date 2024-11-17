package main

import (
	"fmt"
	"os"

	"github.com/rahji/wordsearch"
	"github.com/rahji/wordsearcher/internal/wordlist"
	"github.com/spf13/pflag"
)

func main() {

	// define flags
	var (
		inputFile string
		help      bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
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

	ws := wordsearch.NewWordSearch(16, nil, true)
	unplaced := ws.CreatePuzzle(words)
	if unplaced != nil {
		fmt.Printf("These words could not be placed: %v", unplaced)
	}

	uppercaseGrid := ws.ReturnGrid(wordsearch.GridAllUppercase)
	for i := 0; i < len(uppercaseGrid); i++ {
		for j := 0; j < len(uppercaseGrid[i]); j++ {
			fmt.Printf("%c ", uppercaseGrid[i][j])
		}
		fmt.Println()
	}

	fmt.Println()
	for _, w := range words {
		fmt.Println(w)
	}
}
