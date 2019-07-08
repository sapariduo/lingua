package treebank

import (
	"strings"
	"testing"

	"github.com/sapariduo/lingua"
	"github.com/stretchr/testify/assert"
)

// const sampleConllu = `1	President	President	PROPN	NNP	Number=Sing	2	compound	_	_
// 2	Bush	Bush	PROPN	NNP	Number=Sing	5	nsubj	_	_
// 3	on	on	ADP	IN	_	4	case	_	_
// 4	Tuesday	Tuesday	PROPN	NNP	Number=Sing	5	nmod	_	_
// 5	nominated	nominate	VERB	VBD	Mood=Ind|Tense=Past|VerbForm=Fin	0	root	_	_
// 6	two	two	NUM	CD	NumType=Card	7	nummod	_	_
// 7	individuals	individual	NOUN	NNS	Number=Plur	5	dobj	_	_
// 8	to	to	PART	TO	_	9	mark	_	_
// 9	replace	replace	VERB	VB	VerbForm=Inf	5	advcl	_	_
// 10	retiring	retire	VERB	VBG	VerbForm=Ger	11	amod	_	_
// 11	jurists	jurist	NOUN	NNS	Number=Plur	9	dobj	_	_
// 12	on	on	ADP	IN	_	14	case	_	_
// 13	federal	federal	ADJ	JJ	Degree=Pos	14	amod	_	_
// 14	courts	court	NOUN	NNS	Number=Plur	11	nmod	_	_
// 15	in	in	ADP	IN	_	18	case	_	_
// 16	the	the	DET	DT	Definite=Def|PronType=Art	18	det	_	_
// 17	Washington	Washington	PROPN	NNP	Number=Sing	18	compound	_	_
// 18	area	area	NOUN	NN	Number=Sing	14	nmod	_	_
// 19	.	.	PUNCT	.	_	5	punct	_	_

// `

// sent_id = test-s3
// text = Bagaimana sejarah berdirinya SMAN 1 Pekanbaru yang sudah berusia setengah abad itu?
const sampleConllu = `1	Bagaimana	bagaimana	DET	W--	PronType=Int	2	amod	_	MorphInd=^bagaimana<w>_W--$
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

func Test_ReadConllu(t *testing.T) {
	assert := assert.New(t)
	stu := ReadConllu(strings.NewReader(sampleConllu))
	t.Logf("%v", stu)
	st := ReadConllu(strings.NewReader(sampleConllu))[0]

	// correctHeads := []int{2, 5, 4, 5, 0, 7, 5, 9, 5, 11, 9, 14, 14, 11, 18, 18, 18, 14, 5}
	correctHeads := []int{2, 0, 2, 2, 4, 4, 9, 9, 2, 11, 9, 11, 2}
	assert.Equal(correctHeads, st.Heads)

	// we compare by string to avoid having to build two different test files
	var correctPOS []string
	if lingua.BUILD_TAGSET == "stanfordtags" {
		correctPOS = []string{
			"NNP",
			"NNP",
			"IN",
			"NNP",
			"VBD",
			"CD",
			"NNS",
			"TO",
			"VB",
			"VBG",
			"NNS",
			"IN",
			"JJ",
			"NNS",
			"IN",
			"DT",
			"NNP",
			"NN",
			"FULLSTOP",
		}
	} else {
		correctPOS = []string{
			"DET",
			"NOUN",
			"NOUN",
			"PROPN",
			"PROPN",
			"PROPN",
			"PRON",
			"ADV",
			"VERB",
			"ADP",
			"NOUN",
			"DET",
			"PUNCT",
		}
	}
	t.Logf("%v %T", st.Tags, st.Tags)

	assert.Equal(correctPOS, ttos(st.Tags))
	t.Logf("%v", st.Labels)

	// the stanford tags are not listed in the CONLLU format
	if lingua.BUILD_RELSET != "stanfordrel" {
		var correctRel []string
		correctRel = []string{
			"AMod",
			"Root",
			"Compound",
			"NoDepType",
			"NoDepType",
			"NoDepType",
			"NoDepType",
			"AdvMod",
			"ACl",
			"Case",
			"NoDepType",
			"Det",
			"Punct",
		}

		assert.Equal(correctRel, ltos(st.Labels))
	}
}

func ttos(ts []lingua.POSTag) []string {
	retVal := make([]string, len(ts))
	for i, t := range ts {
		retVal[i] = t.String()
	}
	return retVal
}

func ltos(ls []lingua.DependencyType) []string {
	retVal := make([]string, len(ls))
	for i, l := range ls {
		retVal[i] = l.String()
	}
	return retVal
}
