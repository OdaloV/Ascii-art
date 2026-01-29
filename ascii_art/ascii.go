package ascii_art

func PrintBanner(str string, mapper map[rune][]string) string {
	//Handle  separator i.e '\n'
	if str == "" {
		return ""
	}

strs := []string{}
tempname := ""

for i := 0; i < len(str); i++ {
	if str[i] == '\\' || str[i] == 'n'{
		if tempname!= ""{
			strs = append (strs, tempname)
		}
		if str[i] == 'n'{
			strs = append(strs, "\n")
		}
		tempname = ""
		continue
	}
	tempname += string(str[i])
	if i == len(str) - 1 {
		strs = append(strs,tempname)
	}
}

art_output := ""
 //Start printing characters char by char: by looping into the strs
 for _, word := range strs {
	//Loop into the height which is lines
	for j :=0; j< 8; j++ {
		for i := 0; i < len(word); i++ {
			if word[i] != '\n'{
			x := mapper[rune(word[i])][j]
			if x[0] == '\n' {
				x = x[1:]
			}
			art_output += x
			
		}
		}
		if j == 7 && word == "\n"{
			art_output += "\n"
	}
	if j< 7 && word != "\n"{
		art_output += "\n"
	}
	}
 }
 return art_output + "\n"
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


