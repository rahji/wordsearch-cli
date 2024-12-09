# Wordsearch-cli

Wordsearch-cli is a command-line program that can create a puzzle from a word list -
either from a text file or piped from STDIN. It should work on any computer, but I'm not sure how Windows will like
piping from STDIN. It uses the [Go wordsearch package](http://github.com/rahji/wordsearch) I wrote.

## Install

The easiest way to install this is using `go install`:

```bash
go install github.com/rahji/wordsearch-cli@latest
```

If anyone is interested, I can create Releases so you can just download it for your computer. I don't expect anyone to
use this one though - I will make something more user-friendly.

## Background

Decades ago, I made a website that just allowed people to create Word Search puzzles. I kept it up until recently
because occasionally I would get an email from someone thanking me because it was helping their parent or grandparent
with dementia or Alzheimer's. At some point, it seemed like there were enough other options for creating the puzzles,
so I let the domain name lapse and stopped serving the site.

Now: I've been learning to write programs using Go and I thought this would be a good small project to try and make
happen. I made a package that should allow anyone to make their own wordsearch puzzle generator or interactive game:
http://github.com/rahji/wordsearch

## Usage

Run the command with the `-f` flag to specify the file that contains the word list. Or pipe the list of words to this
command.

`-s` specifies the size of the grid.
