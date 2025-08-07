package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/rahji/wordsearch-cli/internal/wordlist"
	"github.com/rahji/wordsearch/v2"
	"github.com/spf13/pflag"
)

func main() {

	// define flags
	var (
		inputFile    string
		size         int
		sortBy       string
		noOverlap    bool
		noDiagonal   bool
		noBackwards  bool
		showSolution bool
		help         bool
	)

	allowedSortValues := map[string]bool{
		"a-z":  true,
		"z-a":  true,
		"len":  true,
		"rlen": true,
	}
	sortings := make([]string, len(allowedSortValues))
	i := 0
	for k := range allowedSortValues {
		sortings[i] = k
		i++
	}
	slices.Sort(sortings)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.IntVarP(&size, "size", "s", 16, "grid size")
	pflag.StringVarP(&sortBy, "order", "o", "a-z",
		fmt.Sprintf("sorting method for the legend\n%q", sortings),
	)
	pflag.BoolVarP(&noOverlap, "nooverlap", "", false, "disallow overlapping words")
	pflag.BoolVarP(&noDiagonal, "nodiagonal", "", false, "disallow diagonal words")
	pflag.BoolVarP(&noBackwards, "nobackwards", "", false, "disallow backwards-reading words")
	pflag.BoolVarP(&showSolution, "solution", "", false, "show the solution")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	sortBy = strings.ToLower(sortBy)
	_, ok := allowedSortValues[sortBy]
	if !ok {
		fmt.Printf("invalid sorting method: %s. Must be one of %q\n", sortBy, sortings)
		os.Exit(1)
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
	if noBackwards == false {
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
			if grid[i][j] < 97 && showSolution {
				// <97 is uppercase
				bold.Printf("%c ", grid[i][j])
				continue
			}
			if grid[i][j] < 97 {
				fmt.Printf("%c ", grid[i][j])
				continue
			}
			// else lowercase
			fmt.Printf("%c ", grid[i][j]-32)
		}
		fmt.Println()
	}
	fmt.Println()

	switch sortBy {
	case "a-z":
		slices.Sort(words)
	case "z-a":
		slices.Reverse(words)
	case "len":
		sort.Slice(words, func(i, j int) bool {
			return len(words[i]) < len(words[j])
		})
	case "rlen":
		sort.Slice(words, func(i, j int) bool {
			return len(words[i]) > len(words[j])
		})
	}

	for _, w := range words {
		if _, exists := unplacedMap[w]; !exists {
			fmt.Println(w)
		}
	}

	if unplaced != nil {
		fmt.Printf("FYI these %d words could not be placed: %v\n", len(unplaced), unplaced)
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
