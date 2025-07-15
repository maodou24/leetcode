package solution

func isValid(word string) bool {
	var count int
	var vowel, consonant bool
	for _, v := range word {
		if v == '@' || v == '#' || v == '$' {
			return false
		}

		if isChar(v) || isDigit(v) {
			count++
		}
		if isConsonant(v) {
			consonant = true
		}
		if isVowel(v) {
			vowel = true
		}
	}

	return count >= 3 && consonant && vowel
}

var vowel = map[rune]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
}

var consonant = map[rune]struct{}{
	'b': {}, 'c': {}, 'd': {}, 'f': {}, 'g': {}, 'h': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {},
	'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
}

func isChar(v rune) bool {
	return v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z'
}

func isConsonant(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		r = r + 'a' - 'A'
	}
	_, ok := consonant[r]
	return ok
}

func isVowel(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		r = r + 'a' - 'A'
	}
	_, ok := vowel[r]
	return ok
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
