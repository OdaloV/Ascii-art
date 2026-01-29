# ASCII-Art

## Overview

ASCII-Art is a command-line program written in Go that converts a given string into a graphical representation using ASCII characters. Each character is rendered in a large stylized form based on predefined banner templates.

The program supports numbers, letters, spaces, special characters, and newline (`
`) characters, allowing multi-line ASCII-art output.

---

## Objectives

* Read a string from command-line arguments
* Render the string as ASCII-art using banner files
* Handle printable ASCII characters and newline characters
* Preserve input formatting across multiple lines
* Use only standard Go packages
* Follow clean code and good practices

---

## Banner Templates

The project uses banner files to define how each character is displayed.

Available banners:

* `standard`
* `shadow`
* `thinkertoy`

### Banner Format Rules

* Each character has a height of **8 lines**
* Characters are separated by a newline
* Banner files must not be modified

---

## Project Structure

```bash
.
├── ascii_art/
│   ├── ascii_art.go
│   ├── standard.txt
│   ├── shadow.txt
├── main.go
└── Documentation.md
```

---

## Usage

Run the program using `go run` and pass the string as an argument.

Examples:

```
go run . "Hello"
go run . "Hello
There"
go run . "1Hello 2There"
go run . "{Hello There}"
```

The output will be printed directly to standard output in ASCII-art format.

---

## Implementation Overview

### main

* Reads the input string from command-line arguments
* Loads the banner file from the file system
* Builds the ASCII-art mapping
* Prints the rendered output

### Mapper

* Reads a banner file
* Maps each rune to its 8-line ASCII-art representation
* Returns a `map[rune][]string`

### PrintBanner

* Splits input while handling newline characters
* Iterates over each line height (8 lines)
* Renders characters side by side
* Preserves spacing and line breaks

---

## Testing

Unit tests are recommended to verify:

* Correct banner file parsing
* Proper handling of newline characters
* Output height consistency (8 lines per character)
* Edge cases such as empty input

---

## Limitations

* Assumes valid banner files are present
* Expects at least one command-line argument
* Error handling is minimal and focused on file access

---

## Learning Outcomes

This project helps in understanding:

* Go file system
* Data manipulation and mapping
* String and rune handling
* Clean separation of program responsibilities

---

## Author

Flovian Atieno
Faisal Imali
Janet Odalo
---

## License

This project is for learning purpose