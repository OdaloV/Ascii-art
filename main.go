package main
import (
	"ascii/ascii_art"
	"os"
	"fmt"
)

func main() {	
	str := os.Args[1]
	file, err := os.ReadFile("ascii_art/standard.txt")
	if err != nil {
		panic(err)
	}
	mapper := ascii_art.Mapper(file[1:])
	art := ascii_art.PrintBanner(str, mapper)
	fmt.Print(art)
}