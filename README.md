# Advent of Code - Go

Using advent of code to familiarize myself with golang.

The goal is to solve all puzzles without external packages. Forcing myself
to build things like datastructures and algorithms from scratch.

The project structure may change over time as I add solutions and packages.

## Usage

1. Get your inputs for each day on the [official website](https://adventofcode.com/2024)
and add them to the respective `<year>/<day>` directory.
2. `go run ./<year>/<day>/ -input=<path to puzzle input>`
3. The answers will be printed out on the command line

Inputs are in `.gitignore` to not expose them as requested by the Advent of Code guy.

## Personal notes

### I/O

- I/O benchmarking [blog post](https://medium.com/golicious/comparing-ioutil-readfile-and-bufio-scanner-ddd8d6f18463).
Interesting read on bufio.Scanner vs os.ReadFile. 
_Gist: bufio.Scanner is less performant than os.ReadFile but can read line by line._
- To read a whole file it can be better to use os.ReadFile like 
in [this Stack Overflow answer](https://stackoverflow.com/a/66804541).
- `scanner` throws error if line > 65536 characters by default. 
Configue according to [Scanner.Buffer API](https://pkg.go.dev/bufio#Scanner.Buffer).

### Benchmarking & Testing

- [Benchmarking api](https://pkg.go.dev/testing#hdr-Benchmarks).
- Benchmarking funcs measure the whole function, not just the loop. Reset timer 
after expensive set up
- [Blog post](https://medium.com/hyperskill/testing-and-benchmarking-in-go-e33a54b413e) on testing and benchmarking.