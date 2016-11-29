# hashfiles
A tool for calculating the file SHA1 value in a folder, supporting regular expressions.


## Requirements

- Go version >= 1.7


## Installation

To install `hashfiles` use the `go get` command:

```bash
go get github.com/shesuyo/hashfiles
```


## Usage

```
A tool for calculating the file SHA1 value in a folder, supporting regular expressions.

Usage:
	hashfiles command [argument]
The commands are:
	-ignorefile, -if
		the ignore files
	-ignoredir, id
		the ignore dir
	-input, -i
		input dir path
	-output, -o
		out file path
	-goroutine, -g
		the number of goroutine
```