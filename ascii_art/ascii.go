package ascii_art




func GenerateASCIIArt(text string) string {
	// Placeholder implementation
	return "ASCII Art for: " + text
}

func Mapper(file []byte) map[rune][]string {
words := []string{}
word := ""
counter := 1
current_rune := ' '
mapper := make(map[rune][]string)

	for i := 0; i < len(file); i++ {
		if file[i] == '\n' {
			words = append(words, word)
			word = ""
			counter++

		}
		word += string(file[i])

		if counter == 9 {
			mapper[current_rune] = words
			if i != len(file) - 1 {
				i++
			}

			counter = 1
			words = nil
			current_rune++
		}
	} 


	return mapper	

}


