package wordlist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// GetWords reads from an input filename or from STDIN if the filename is empty.
// It returns a slice of words from the input, after removing any non-letter characters.
func GetWords(infile string) ([]string, error) {

	var reader io.Reader

	// if file flag is provided, try to open that file
	if infile != "" {
		file, err := os.Open(infile)
		if err != nil {
			return nil, fmt.Errorf("Error opening file: %v", err)
		}
		defer file.Close()
		reader = file
	} else {
		// get the word list from STDIN
		stat, _ := os.Stdin.Stat()
		// unless the STDIN is not piped data
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			errString := "No input file specified and no data piped to STDIN\n" +
				"Usage: either pipe data to STDIN or use -f flag to specify input file"
			return nil, fmt.Errorf(errString)
		}
		reader = os.Stdin
	}

	// create a scanner to read the input, line by line
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	// turn each line/string into a "word" and add it to the slice of words
	var words []string
	for scanner.Scan() {
		line := scanner.Text()
		word := cleanString(line)

		// if the line is empty now, skip it
		if word == "" {
			continue
		}
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading input: %v", err)
	}

	return words, nil
}

// Remove non-letter characters from a string
func cleanString(s string) (ret string) {
	re, _ := regexp.Compile("[^[:alpha:]]")
	replacement := ""
	ret = string(re.ReplaceAll([]byte(s), []byte(replacement)))
	return
}
