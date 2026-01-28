package main
import (
	"ascii/ascii_art"
	"os"
	"fmt"
)

func main() {	// This is a placeholder for the main function.
	str := os.Args[1]
	file, err := os.ReadFile("ascii_art/standard.txt")
	if err != nil {
		panic(err)
	}
	mapper := ascii_art.Mapper(file[1:])
	ascii_art.PrintBanner(str, mapper)
	fmt.Println()
}