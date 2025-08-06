package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rahji/wordsearch-cli/internal/wordlist"
	"github.com/rahji/wordsearch/v2"
	"github.com/spf13/pflag"
)

func main() {

	// define flags
	var (
		inputFile  string
		size       int
		noOverlap  bool
		noDiagonal bool
		noReverse  bool
		help       bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.IntVarP(&size, "size", "s", 16, "grid size")
	pflag.BoolVarP(&noOverlap, "nooverlap", "o", false, "disallow overlapping words")
	pflag.BoolVarP(&noDiagonal, "nodiagonal", "d", false, "disallow diagonal words")
	pflag.BoolVarP(&noReverse, "noreverse", "r", false, "disallow reverse words")
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

	directions := []string{"S", "E"}
	if noDiagonal == false {
		directions = append(directions, "SE", "SW")
	}
	if noReverse == false {
		directions = append(directions, reverseCardinals(directions)...)
	}

	var ws *wordsearch.WordSearch
	if noOverlap {
		ws = wordsearch.NewWordSearch(size,
			wordsearch.WithDirections(directions),
			wordsearch.WithoutOverlaps(),
		)
	} else {
		ws = wordsearch.NewWordSearch(size,
			wordsearch.WithDirections(directions),
		)
	}

	unplaced := ws.CreatePuzzle(words)

	// make a map of the unplaced words for later lookup
	unplacedMap := make(map[string]struct{})
	for _, word := range unplaced {
		unplacedMap[word] = struct{}{}
	}

	bold := color.New(color.FgBlue)

	grid := ws.ReturnGrid(wordsearch.GridRaw)
	for i := range grid {
		for j := range grid[i] {
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

// reverseCardinals accepts a slice of unique cardinal direction abbreviations
// and returns a slice containing the reverse directions for each of them
func reverseCardinals(dirs []string) []string {
	var ret []string
	for _, d := range dirs {
		switch d {
		case "S":
			ret = append(ret, "N")
		case "E":
			ret = append(ret, "W")
		case "W":
			ret = append(ret, "E")
		case "N":
			ret = append(ret, "S")
		case "SE":
			ret = append(ret, "NW")
		case "SW":
			ret = append(ret, "NE")
		case "NE":
			ret = append(ret, "S")
		case "NW":
			ret = append(ret, "S")
		}
	}
	return ret
}
