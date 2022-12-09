# adventofcode-fetcher
CLI command that fetches inputs for Advent Of Code event puzzles

## Installation
### go install
```
go install github.com/streakinthesky/adventofcode-fetcher/aoc
```

## To Use
Log into (Advent Of Code)[https://adventofcode.com] and grab the value of the cookie named `session` with your browser's dev tools([for chrome](https://developer.chrome.com/docs/devtools/storage/cookies/)).

You can copy and paste the value in a file called `session` in the current directory. It must be the raw value on the first line of the file. Then use the `aoc` command with the url of the puzzle inputs you want to fetch:
```
aoc fetch https://adventofcode.com/2022/day/1
```

OR

You can pass it directly into the command using the `-session` flag (ie. the session value is 1234567890abcdef)
```
aoc fetch --session=1234567890abcdef https://adventofcode.com/2022/day/1
```

OR

You can save it in any file and pass the file path into the `-session` flag (ie. saved in a file `../path/to/file`)
```
aoc fetch --session ../path/to/file https://adventofcode.com/2022/day/1
```

These will save the inputs for the event in a file called `inputs.txt` in the current directory.
