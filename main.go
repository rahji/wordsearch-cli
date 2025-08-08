package main

// get multicol to work in text with spaces
// fix left align for legend

import (
	"embed"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/rahji/wordsearch/v2"
	"github.com/rahji/wordsearcher/internal/wordlist"
	"github.com/signintech/gopdf"
	"github.com/spf13/pflag"
)

//go:embed RobotoMono-Regular.ttf
var embeddedFile embed.FS

// define flags
var (
	inputFile    string
	outputFile   string
	title        string
	showSolution bool
	size         int
	sortBy       string
	cols         int
	noOverlap    bool
	noDiagonal   bool
	noBackwards  bool
	help         bool
)

const (
	fontSize          = 14
	topMargin         = 20.0
	gridLineSpacing   = 24.0
	legendLineSpacing = 20.0
	titleSpacing      = 40.0
	colSpacing        = 5
	gridCharSpacing   = 4.0
	normalCharSpacing = 1.0
)

func main() {

	parseFlags()

	pdf := gopdf.GoPdf{}
	if outputFile != "" {
		file, err := embeddedFile.Open("RobotoMono-Regular.ttf")
		if err != nil {
			fmt.Printf("couldn't open font file: %s", err.Error())
			os.Exit(1)
		}
		defer file.Close()

		fontContainer := &gopdf.FontContainer{}
		err = fontContainer.AddTTFFontByReader("Roboto", file)
		if err != nil {
			fmt.Printf("couldn't make font container: %s", err.Error())
			os.Exit(1)
		}
		pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeLetter})
		pdf.SetMarginTop(topMargin)
		pdf.AddPage()
		err = pdf.AddTTFFontFromFontContainer("Roboto", fontContainer)
		if err != nil {
			fmt.Printf("couldn't add font from container: %s", err.Error())
			os.Exit(1)
		}
		err = pdf.SetFont("Roboto", "", fontSize)
		if err != nil {
			fmt.Printf("couldn't set the font: %s", err.Error())
			os.Exit(1)
		}
		if title != "" {
			pdf.CellWithOption(
				&gopdf.Rect{W: gopdf.PageSizeLetter.W},
				title,
				gopdf.CellOption{Align: gopdf.Center},
			)
			pdf.Br(titleSpacing)
		}
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

	if title != "" {
		fmt.Printf("%s\n\n", title)
	}
	grid := ws.ReturnGrid(wordsearch.GridRaw)
	gridText := formattedGrid(grid)
	fmt.Println(gridText)
	fmt.Println()
	if outputFile != "" {
		pdf.SetCharSpacing(gridCharSpacing)
		textToPDF(&pdf,
			gridText,
			gridLineSpacing,
			gopdf.PageSizeLetter.W,
			gopdf.CellOption{Align: gopdf.Center},
		)
	}

	// filter out unplaced words before writing legend
	n := 0
	for _, w := range words {
		if _, exists := unplacedMap[w]; !exists {
			words[n] = w
			n++
		}
	}
	words = words[:n]

	legendText := formattedLegend(words, cols, sortBy)
	fmt.Println(legendText)
	if outputFile != "" {
		pdf.SetCharSpacing(normalCharSpacing)
		textToPDF(&pdf,
			legendText,
			legendLineSpacing,
			gopdf.PageSizeLetter.W,
			gopdf.CellOption{Align: gopdf.Center},
		)
	}

	if unplaced != nil {
		fmt.Printf("\nWARNING: These %d words could not be placed: %v\n", len(unplaced), unplaced)
	}

	// things could definitely be refactored, with a seperate function that outputs a title, grid, and legend
	// but adding the solution was kind of an afterthought and it ended up this way, which is meh

	if showSolution {
		fmt.Printf("\n%s (SOLUTION)\n\n", title)
		solutionGrid := ws.ReturnGrid(wordsearch.GridWithDots)
		gridText = formattedGrid(solutionGrid)
		fmt.Println(gridText)
		if outputFile != "" {
			pdf.AddPage()
			pdf.CellWithOption(
				&gopdf.Rect{W: gopdf.PageSizeLetter.W},
				title+" (SOLUTION)",
				gopdf.CellOption{Align: gopdf.Center},
			)
			pdf.Br(titleSpacing)
			pdf.SetCharSpacing(gridCharSpacing)
			textToPDF(&pdf,
				gridText,
				gridLineSpacing,
				gopdf.PageSizeLetter.W,
				gopdf.CellOption{Align: gopdf.Center},
			)
		}
	}

	if outputFile != "" {
		pdf.WritePdf(outputFile)
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
		if len(w) > longest {
			longest = len(w)
		}
	}
	return longest
}

// formattedGrid returns the grid as a printable string,
// (remember that the grid uses uppercase for letters that are part of the solution,
// and lowercase letters for "filler" letters)
func formattedGrid(grid [][]byte) string {
	ret := strings.Builder{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] < 97 { // uppercase
				ret.WriteByte(grid[i][j])
				ret.WriteString(" ")
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

func formattedLegend(words []string, cols int, sortBy string) string {
	var ret strings.Builder

	switch sortBy {
	case "a-z":
		slices.Sort(words)
	case "z-a":
		slices.Sort(words)
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

	longest := longestLength(words)
	half := int(math.Ceil(float64(len(words)) / float64(cols)))
	for i := range half {
		for c := range cols {
			ic := i + half*c // index for this column's entry
			if ic > len(words)-1 {
				// add empty space (without padding) for the last entry, to make centering work
				ret.WriteString(strings.Repeat(" ", longest))
				break
			}
			padding := longest - len(words[ic])
			if c < cols-1 {
				padding += colSpacing
			}
			ret.WriteString(words[ic])
			ret.WriteString(strings.Repeat(" ", padding))
		}
		ret.WriteString("\n")
	}

	return ret.String()
}

func parseFlags() {
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
	pflag.StringVarP(&outputFile, "pdf", "", "", "output PDF file (instead of STDOUT)")
	pflag.StringVarP(&title, "title", "", "", "title of the puzzle")
	pflag.BoolVarP(&showSolution, "solution", "", false, "also show the solution")
	pflag.StringVarP(&sortBy, "order", "o", "a-z",
		fmt.Sprintf("sorting method for the legend\n%q", sortings),
	)
	pflag.IntVarP(&cols, "cols", "", 2, "legend columns")
	pflag.BoolVarP(&noOverlap, "nooverlap", "", false, "disallow overlapping words")
	pflag.BoolVarP(&noDiagonal, "nodiagonal", "", false, "disallow diagonal words")
	pflag.BoolVarP(&noBackwards, "nobackwards", "", false, "disallow backwards-reading words")
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
}

// textToPDF replaces newlines with pdf.Br() using a lineSpacing argument
func textToPDF(pdf *gopdf.GoPdf, text string, spacing float64, width float64, opt gopdf.CellOption) {
	lines := strings.SplitSeq(text, "\n")
	for l := range lines {
		pdf.CellWithOption(&gopdf.Rect{W: width}, l, opt)
		pdf.Br(spacing)
	}
}
