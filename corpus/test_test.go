package corpus

import (
	"strings"

	"github.com/sapariduo/lingua/treebank"
)

func mediumSentence() []treebank.SentenceTag {
	conllu := `1	Bagaimana	bagaimana	DET	W--	PronType=Int	2	amod	_	MorphInd=^bagaimana<w>_W--$
2	sejarah	sejarah	NOUN	NSD	Number=Sing	0	root	_	MorphInd=^sejarah<n>_NSD$
3	berdirinya	berdiri	NOUN	VSA+PS3	Number=Sing|Number[psor]=Sing|Person[psor]=3|Voice=Act	2	compound	_	MorphInd=^ber+diri<n>_VSA+dia<p>_PS3$
4	SMAN	sman	PROPN	X--	_	2	flat	_	MorphInd=^sman<x>_X--$
5	1	1	PROPN	CC-	_	4	flat	_	MorphInd=^1<c>_CC-$
6	Pekanbaru	pekanbaru	PROPN	X--	_	4	flat	_	MorphInd=^pekanbaru<x>_X--$
7	yang	yang	PRON	S--	PronType=Rel	9	nsubj:pass	_	MorphInd=^yang<s>_S--$
8	sudah	sudah	ADV	D--	_	9	advmod	_	MorphInd=^sudah<d>_D--$
9	berusia	berusia	VERB	VSA	Number=Sing|Voice=Act	2	acl	_	MorphInd=^ber+usia<n>_VSA$
10	setengah	setengah	ADP	NSD	Number=Sing	11	case	_	MorphInd=^setengah<n>_NSD$
11	abad	abad	NOUN	NSD	Number=Sing	9	obl	_	MorphInd=^abad<n>_NSD$
12	itu	itu	DET	B--	PronType=Dem	11	det	_	SpaceAfter=No|MorphInd=^itu<b>_B--$
13	?	?	PUNCT	Z--	_	2	punct	_	MorphInd=^?<z>_Z--$	

`
	readr := strings.NewReader(conllu)
	return treebank.ReadConllu(readr)
}

const EPSILON64 float64 = 1e-10

func floatEquals64(a, b float64) bool {
	if (a-b) < EPSILON64 && (b-a) < EPSILON64 {
		return true
	}
	return false
}
