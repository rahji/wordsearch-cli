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
Showing a grid using `wordsearch.GridWithSpaces`:

```
O S             O B I W A N
L R I B O N E K
O A       S K Y W A L K E R
S W     C H E W B A C C A
N R             H T I S
A A       D E A T H S T A R
H T             E M P I R E
  S   A   T A T O O I N E
A     D N       I
D     R   A       D
O L   O     K       E
Y U   I       I       J
  K   D         N
  E
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

## Usage

Run the command with the `-f` flag to specify the file that contains the word list. Or pipe the list of words to this
command.

`-s` specifies the size of the grid.
