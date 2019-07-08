// +build !stanfordtags

package treebank

import "github.com/sapariduo/lingua"

var posTagTable map[string]lingua.POSTag = map[string]lingua.POSTag{
	"X":     lingua.X,
	"ADJ":   lingua.ADJ,
	"ADP":   lingua.ADP,
	"ADV":   lingua.ADV,
	"AUX":   lingua.AUX,
	"CONJ":  lingua.CONJ,
	"DET":   lingua.DET,
	"INTJ":  lingua.INTJ,
	"NOUN":  lingua.NOUN,
	"NUM":   lingua.NUM,
	"PART":  lingua.PART,
	"PRON":  lingua.PRON,
	"PROPN": lingua.PROPN,
	"PUNCT": lingua.PUNCT,
	"SCONJ": lingua.SCONJ,
	"SYM":   lingua.SYM,
	"VERB":  lingua.VERB,

	"-NULL-":    lingua.X,
	"-ROOT-":    lingua.ROOT_TAG,
	"-UNKNOWN-": lingua.UNKNOWN_TAG,
}
