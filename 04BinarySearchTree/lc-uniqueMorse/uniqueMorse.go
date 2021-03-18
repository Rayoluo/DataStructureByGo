package lc_uniqueMorse

func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]struct{})
	table := [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for i := 0; i < len(words); i++ {
		str := ""
		for j := 0; j < len(words[i]); j++ {
			str += table[words[i][j]-'a']
		}
		set[str] = struct{}{}
	}
	return len(set)
}
