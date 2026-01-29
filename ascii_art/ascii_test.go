package ascii_art

import (
	"testing"
	"strconv"
	
)

var mapp = createInputMoker()

var tests = []struct {
		name string
		input    string
		expected string
	}{ 
		{"Test empty string", "", createOutputMoker("", mapp)},
		{"Test single character", "A", createOutputMoker("A", mapp)},
		{"Test special characters", "@#&*", createOutputMoker("@#&*", mapp)},
		{"Test numbers", "1234567890", createOutputMoker("1234567890", mapp)},
		{"test with newline", "Hello\\nWorld!", createOutputMoker("Hello\\nWorld!", mapp)},
		{"Test multiple lines", "Line1\\nLine2\\nLine3", createOutputMoker("Line1\\nLine2\\nLine3", mapp)},
		{"test with space", "Hello World", createOutputMoker("Hello World", mapp)},
		{"test all ascii", " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", createOutputMoker(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", mapp)},
	
	}

func createInputMoker() map[rune][]string {
	// we will create mock data for testing, instead of reading from file, we will create for all ascii and instead of having the art as real we will have simple representation ie. for 'a' we will have ["a1","a2",...,"a8"]
	mapper := make(map[rune][]string)
	for i := 32; i <= 126; i++ {
		char := rune(i)
		lines := []string{}
		for j := 1; j <= 8; j++ {
			lines = append(lines, string(char)+string(strconv.Itoa(j)))
		}
		mapper[char] = lines
	}
	return mapper
}

func createOutputMoker(str string, mapper map[rune][]string) string {
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



func TestPrintBanner(t *testing.T) {
	mapper := createInputMoker()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := PrintBanner(tt.input, mapper)
			if result != tt.expected {
				t.Errorf("PrintBanner(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMapper(t *testing.T) {
	// Fake ASCII-art data for 2 characters: ' ' and '!'
	// Each character has 8 lines, separated by '\n'
	file := []byte(
		" \n \n \n \n \n \n \n \n" + // character ' '
			"!\n!\n!\n!\n!\n!\n!\n!\n", // character '!'
	)

	expected := map[rune][]string{
		' ': {" ", " ", " ", " ", " ", " ", " ", " "},
		'!': {"!", "!", "!", "!", "!", "!", "!", "!"},
	}

	result := Mapper(file)

	if result[' '] == nil || result['!'] == nil {
		t.Errorf("Mapper() = %v; want %v", result, expected)
	}
}

	