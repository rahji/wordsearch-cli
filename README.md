# Wordsearcher

Decades ago, I made a website that just allowed people to create Word Search puzzles. I kept it up until recently
because occasionally I would get an email from someone thanking me because it was helping their parent or grandparent
with dementia or Alzheimer's. At some point, it seemed like there were enough other options for creating the puzzles,
so I let the domain name lapse and stopped serving the site.

Now: I've been learning to write programs using Go and I thought this would be a good small project to try and make
happen. The package I made is here: http://github.com/rahji/wordsearch

This repo is a CLI program that can create a puzzle from a word list - either from a text file or piped from STDIN.

I'm trying to figure out what the next iteration will be. These are the options I'm considering:

- A simple CLI application with minimal options. (That's this repo)
- A fancier TUI application using Bubble Tea. This would still require access via a terminal, but it would afford
  feedback while making the puzzle. You could regenerate the whole puzzle if you wanted a new random arrangement of the
  same word list, or you could add a word while keeping the current puzzle layout, for example. I could also make it
  accessible via SSH using Wish, which would not require any kind of installation but would limit the types of output.
- A GUI application that uses Fyne. This would be nice because it wouldn't require someone to know anything about the
  command-line, it would work on any operating system, and would again allow some instant feedback when making the
  puzzle.

None of the options would require internet access (except for the SSH version). It would be nice if any/all of the
options could output text or a PDF file (or maybe an HTML file).