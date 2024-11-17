# Wordsearch-cli

Wordsearch-cli is a command-line program that can create a puzzle from a word list -
either from a text file or piped from STDIN. It should work on any computer, but I'm not sure how Windows will like
piping from STDIN. It uses the [Go wordsearch package](http://github.com/rahji/wordsearch) I wrote.

## Example

```
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

## Next Steps

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
