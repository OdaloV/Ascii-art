package ascii_art

func GenerateASCIIArt(text string) string {
	// Placeholder implementation
	return "ASCII Art for: " + text
}
package ascii_art

func GenerateASCIIArt(text string) string {
	// Placeholder implementation
	return "ASCII Art for: " + text
}
func PrintBanner(s string, mapper map) {
 //Start printing characters char by char: by looping into the strs
 for _, word := range strs {
	//Loop into the height which is lines
	for j :=0: j< 8; j++ {
		for i := 0; i < len(word); i++ {
			if word[i] != '\n'
			x := mapper[rune(word[i])][j]
			if x[0] == '\n' {
				x = x[1:]
			}
			fmt.Print(x)
		}
	}
	if j==7 && word == "\n"{
		fmt.Println()
	}
	if j< 7 && word != "\n"{
		fmt.Print
	}
 }
}