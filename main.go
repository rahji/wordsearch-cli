package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

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

	grid := ws.ReturnGrid(wordsearch.GridRaw)
	fmt.Println(formattedGrid(grid, false))
	fmt.Println()

	// filter out unplaced words
	n := 0
	for _, w := range words {
		if _, exists := unplacedMap[w]; !exists {
			words[n] = w
			n++
		}
	}
	words = words[:n]

	// sort them before displaying the legend
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

	// longest := longestLength(words)
	// cols := ((size*2)-1 > (longest*2)+3) // true means there is room for 2 columns
	// xxx filter slice based on unplaced map
	// if !cols print if it's not in the unplaced map
	// if cols loop from 1 to length of words - unplaced length
	// if cols {
	// 	for i:= 0; i<=len(words)/2; i++ {
	// 		if _, exists := unplacedMap[words[i]]; !exists {
	// 			fmt.Printf("%"+longest+"s   "%s\n", w)
	// 		}
	// 	}
	// } else {}

	for _, w := range words {
		fmt.Println(w)
	}

	if unplaced != nil {
		fmt.Printf("\nWARNING: These %d words could not be placed: %v\n", len(unplaced), unplaced)
	}

	if showSolution {
		fmt.Print("\n\n")
		fmt.Println(formattedGrid(grid, true))
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

// longestLength returns the number of characters in a slice's longest word
func longestLength(words []string) int {
	longest := 0
	for _, w := range words {
		if len(w) > 0 {
			longest = len(w)
		}
	}
	return longest
}

// formattedGrid returns the grid (or the solution) as a printable string,
// (remember that the grid uses uppercase for letters that are part of the solution,
//
//	and lowercase letters for "filler" letters)
func formattedGrid(grid [][]byte, solution bool) string {
	ret := strings.Builder{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] < 97 { // uppercase
				ret.WriteByte(grid[i][j])
				ret.WriteString(" ")
				continue
			}
			if solution { // lowercase, but we want the solution only
				ret.WriteString("  ")
				continue
			}
			// lowercase, but show it as uppercase
			ret.WriteByte(grid[i][j] - 32)
			ret.WriteString(" ")
		}
		ret.WriteString("\n")
	}
	return ret.String()
}
