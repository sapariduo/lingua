package treebank

import "github.com/sapariduo/lingua"

var alreadyLogged map[string]bool = make(map[string]bool)

// TODO : CHECK
func StringToLexType(tag string) lingua.LexemeType {
	var lexType lingua.LexemeType
	switch tag {
	case "NUM":
		lexType = lingua.Number
	case "PUNCT":
		lexType = lingua.Punctuation
	case "SYM":
		lexType = lingua.Symbol
	default:
		lexType = lingua.Word
	}
	return lexType
}

func StringToPOSTag(tag string) (lingua.POSTag, bool) {
	t, ok := posTagTable[tag]

	return t, ok
}

func StringToDependencyType(ud string) (lingua.DependencyType, bool) {
	dt, ok := dependencyTable[ud]

	return dt, ok
}

func reset() (lingua.LexemeSentence, []lingua.POSTag, []int, []lingua.DependencyType) {
	s := lingua.NewLexemeSentence()
	st := make([]lingua.POSTag, 0)
	sh := make([]int, 0)
	sdt := make([]lingua.DependencyType, 0)

	return s, st, sh, sdt
}

func finish(s lingua.LexemeSentence, st []lingua.POSTag, sh []int, sdt []lingua.DependencyType, sentences []SentenceTag) []SentenceTag {
	sentenceTag := SentenceTag{s, st, sh, sdt}
	sentences = append(sentences, sentenceTag)

	return sentences
}
