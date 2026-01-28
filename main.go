package main
import (
	"ascii/ascii_art"
	"os"
)

func main() {	// This is a placeholder for the main function.
	str := os.Args[1]
	file, err := os.ReadFile("ascii_art/mapping.txt")
	if err != nil {
		panic(err)
	}
	mapper := ascii_art.Mapper(file)
	ascii_art.PrintBanner(str, mapper)
}