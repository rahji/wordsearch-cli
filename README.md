# Wordsearch-cli

Wordsearch-cli is a command-line program that can create a puzzle from a word list -
either from a text file or piped from STDIN. It produces a plaintext version of the output and
optionally makes a PDF file, with some options for how it's formatted. It can also produce
the solution at the same time.

This should work on any computer, but I'm not sure how Windows will like
piping from STDIN. It uses the [Go wordsearch package](http://github.com/rahji/wordsearch) I wrote.

## Example

```
STAR WARS

W C Z K B G F B B E U L S S P T
S I F B E M C O F B V U T S I F
N D O V M N N K F Z K K R E M M
I E A U J D O A Y F M E L Y J R
K J H C H E W B A C C A H I V X
A G J R N E C L I Z H N A A N O
N R N E U R E K L A W Y K S D R
A S A X O B I W A N O K V A S R
Z E I T G F Y E N I O O T A T X
K M N T S U C O F E W U D M E R
P P F F H H M H D E E P N R N O
G I Q Y O A T A S A G Z W D J H
O R S R A W R A T S T X R L U X
Y E V M M V Q B E A T O J S L D
R H A N S O L O K D I S O X U K
O L A U K I W C Q D L Q P Y Y I

DEATHSTAR   SKYWALKER
STARWARS    HANSOLO
EMPIRE      ANAKIN
DROID       SITH
OBIWAN      KENOBI
CHEWBACCA   TATOOINE
JEDI        LUKE
YODA
```

## Install

Download the appropriate binary from the Releases page and place it somewhere in your PATH.
Then type `wordsearch-cli` and some options from below in your terminal application.

For example:

```bash
./wordsearch-cli --cols 3 --size 16 -f starwars.txt --pdf puzzle.pdf --title "STAR WARS" --solution
```

(Of course, if you have Go installed, you can install by typing: this is using `go install github.com/rahji/wordsearch-cli@latest`)

## Background

Decades ago, I made a website that just allowed people to create Word Search puzzles. I kept it up until recently
because occasionally I would get an email from someone thanking me because it was helping their parent or grandparent
with dementia or Alzheimer's. At some point, it seemed like there were enough other options for creating the puzzles,
so I let the domain name lapse and stopped serving the site.

Now: I've been learning to write programs using Go and I thought this would be a good small project to try and make
happen. I made a package that should allow anyone to make their own wordsearch puzzle generator or interactive game:
http://github.com/rahji/wordsearch/v2

## Usage

Run the command with the `-f` flag to specify the file that contains the word list. Or pipe the list of words to this
command.

The rest of the options:

```
Usage of ./wordsearch-cli:
  -f, --file string    input file (if not specified, reads from STDIN)
  -s, --size int       grid size (default 16)
      --pdf string     output PDF file (instead of STDOUT)
      --title string   title of the puzzle
      --solution       also show the solution
  -o, --order string   sorting method for the legend
                       ["a-z" "len" "rlen" "z-a"] (default "a-z")
      --cols int       legend columns (default 2)
      --nooverlap      disallow overlapping words
      --nodiagonal     disallow diagonal words
      --nobackwards    disallow backwards-reading words
  -h, --help           show help message
```

> TIP: to make an easy puzzle, you might use the `--nodiagonal` and `--nobackwards` options with a small `--size`
