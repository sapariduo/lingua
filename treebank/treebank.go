package treebank

import (
	"archive/zip"
	"io"

	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sapariduo/lingua"
)

var empty struct{}

// Loader is anything that loads into a slice of SentenceTags. For future uses, to load tree banks
type Loader func(string) []SentenceTag

// LoadUniversal loads a treebank file formatted in a CONLLU format
func LoadUniversal(fileName string) []SentenceTag {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return ReadConllu(f)
}

// ReadConllu reads a file formatted in a CONLLU format
func ReadConllu(reader io.Reader) []SentenceTag {
	s, st, sh, sdt := reset()
	sentences := make([]SentenceTag, 0)
	sentenceCount := 0

	var usedTags lingua.TagSet
	var usedDepTypes lingua.DependencyTypeSet
	var unknownTags = make(map[string]struct{})
	var unknownDepType = make(map[string]struct{})

	colCount := 0
	for bs := bufio.NewScanner(reader); bs.Scan(); colCount++ {
		l := bs.Text()

		if len(l) == 0 {
			// then this is a new sentence
			sentences = finish(s, st, sh, sdt, sentences)
			s, st, sh, sdt = reset()

			sentenceCount++
			continue
		}

		cols := strings.Split(l, "\t")
		word := cols[1]

		var tag string
		switch lingua.BUILD_TAGSET {
		case "stanfordtags":
			tag = cols[4]
		case "universaltags":
			tag = cols[3]
		default:
			panic("Unknown tagset")
		}

		head := cols[6]
		depType := cols[7]

		var t lingua.POSTag
		var dt lingua.DependencyType
		var h int
		var ok bool
		var err error

		word = lingua.UnescapeSpecials(word)

		lexType := StringToLexType(tag)
		if t, ok = StringToPOSTag(tag); ok {
			usedTags[t] = true
		} else {
			unknownTags[tag] = empty
		}

		if h, err = strconv.Atoi(head); err != nil {
			panic(err) // panic is the right option, because there is no default
		}

		if dt, ok = StringToDependencyType(depType); ok {
			usedDepTypes[dt] = true
		} else {
			unknownDepType[depType] = empty
		}

		lexeme := lingua.Lexeme{word, lexType, sentenceCount, colCount}
		s = append(s, lexeme)
		st = append(st, t)
		sh = append(sh, h)
		sdt = append(sdt, dt)
	}
	return sentences
}

// LoadEWT loads a zipped English Web Treebank (as donated by Google)
func LoadEWT(filename string) []SentenceTag {

	r, err := zip.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	sentences := make([]SentenceTag, 0)

	for _, f := range r.File {
		contents, err := f.Open()
		if err != nil {
			panic(err)
		}
		sentences = append(sentences, ReadConllu(contents)...)
		contents.Close()
	}

	return sentences
}
