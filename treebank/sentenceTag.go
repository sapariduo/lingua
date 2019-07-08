package treebank

import (
	"math/rand"

	"github.com/sapariduo/lingua"
)

// SentenceTag is a struc that holds a sentence, tags, heads and labels
type SentenceTag struct {
	Sentence lingua.LexemeSentence
	Tags     []lingua.POSTag
	Heads    []int
	Labels   []lingua.DependencyType
}

func (s SentenceTag) AnnotatedSentence(f lingua.AnnotationFixer) lingua.AnnotatedSentence {
	retVal := lingua.NewAnnotatedSentence()
	retVal = append(retVal, lingua.RootAnnotation())

	for i, lex := range s.Sentence {
		a := lingua.NewAnnotation()
		a.Lexeme = lex
		a.POSTag = s.Tags[i]
		a.DependencyType = s.Labels[i]

		// should panic, because SentenceTag is only ever used during training
		if err := a.Process(f); err != nil {
			panic(err)
		}

		retVal = append(retVal, a)
	}

	// add heads
	for i, a := range retVal {
		if i == 0 {
			continue
		}
		a.SetHead(retVal[s.Heads[i-1]])
	}

	retVal.Fix()

	return retVal
}

func (s SentenceTag) Dependency(f lingua.AnnotationFixer) *lingua.Dependency {
	sentence := s.AnnotatedSentence(f)
	dep := sentence.Dependency()

	return dep
}

func (s SentenceTag) String() string {
	return s.Sentence.String()
}

func ShuffleSentenceTag(s []SentenceTag) []SentenceTag {
	rand.Seed(1337)
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	return s
}

/* UTILITY FUNCTIONS */

func WrapLexemeSentence(sentence lingua.LexemeSentence) lingua.LexemeSentence {
	retSentence := lingua.NewLexemeSentence()
	retSentence = append(retSentence, lingua.StartLexeme())
	retSentence = append(retSentence, sentence...)
	retSentence = append(retSentence, lingua.RootLexeme())
	return retSentence
}

func WrapTags(tagList []lingua.POSTag) []lingua.POSTag {
	retVal := append([]lingua.POSTag{lingua.X}, tagList...)
	retVal = append(retVal, lingua.X)
	return retVal
}

func WrapHeads(heads []int) []int {
	retVal := append([]int{0}, heads...)
	retVal = append(retVal, 0)
	return retVal
}

func WrapDeps(deps []lingua.DependencyType) []lingua.DependencyType {
	retVal := append([]lingua.DependencyType{lingua.Dep}, deps...)
	retVal = append(retVal, lingua.Dep)
	return retVal
}
