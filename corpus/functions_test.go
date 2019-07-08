package corpus

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateCorpus(t *testing.T) {
	sentenceTags := mediumSentence()
	dict := GenerateCorpus(sentenceTags)

	// testing time
	assert := assert.New(t)
	expectedWords := []string{"", "-UNKNOWN-", "-ROOT-", "Bagaimana", "sejarah", "berdirinya", "SMAN", "1", "Pekanbaru", "yang", "sudah", "berusia", "setengah", "abad", "itu", "?"}

	expectedIDs := make(map[string]int)
	for i, w := range expectedWords {
		expectedIDs[w] = i
	}

	assert.Equal(expectedWords, dict.words, "Corpus known words should be the same as the manually annotated expected values")
	assert.Equal(expectedIDs, dict.ids, "IDs should be the same as expected IDs")
	assert.Equal(int64(len(expectedWords)), dict.maxid)
}

func TestViterbiSplit(t *testing.T) {
	assert := assert.New(t)
	dict := GenerateCorpus(mediumSentence())

	s2 := "sejarahberdirinya"
	words := ViterbiSplit(s2, dict)
	assert.Equal([]string{"sejarah", "berdirinya"}, words)

	s2 = "yangsudahberusia"
	words = ViterbiSplit(s2, dict)
	assert.Equal([]string{"yang", "sudah", "berusia"}, words)

	s3 := "setengahabad"
	words = ViterbiSplit(s3, dict)
	assert.Equal([]string{"setengah", "abad"}, words)
}

func TestCosineSimilarity(t *testing.T) {
	a := strings.Split("Membeli roti di kota tua", " ")
	b := strings.Split("Membeli sebuah roti di kota tua", " ")

	s1 := CosineSimilarity(a, a)
	s2 := CosineSimilarity(a, b)

	if !floatEquals64(s1, 1) {
		t.Error("Expected similarity to be 1 when compared with itself")
	}
	if s2 > s1 {
		t.Error("Something went wrong with the cosine similarity algorithm")
	}

	c := strings.Split("Jalan Pondok Labu", " ")
	d := strings.Split("Jln Pondok Labu Raya", " ")

	s1 = CosineSimilarity(c, c)
	s2 = CosineSimilarity(c, d)
	t.Logf("cosine similarity value: %v\n", s2)

	if !floatEquals64(s1, 1) {
		t.Error("Expected similarity to be 1 when compared with itself")
	}
	if s2 > s1 {
		t.Error("Something went wrong with the cosine similarity algorithm")
	}
}

func TestDL(t *testing.T) {
	a := "TELEKOMUNIKASI INDONESIA"
	b := "TELEKOMUNIKASI INDONESIA INTERNASIONAL"

	s1 := DamerauLevenshtein(a, a)
	s2 := DamerauLevenshtein(a, b)

	if s1 != 0 {
		t.Errorf("Expected the distance to be 0 when compared against itself. Got %d", s1)
	}

	if s2 < s1 {
		t.Error("Expected DL similarity to be greater when compared against itself")
	}

	c := "Pondok Labu"
	d := "Jl Pondok Labu"
	e := "Pondok Labu Ujung"

	s1 = DamerauLevenshtein(c, c)
	s2 = DamerauLevenshtein(c, d)
	s3 := DamerauLevenshtein(c, e)
	t.Logf("DL value: %v", s2)
	t.Logf("DL value: %v", s3)

	if s1 != 0 {
		t.Errorf("Expected the distance to be 0 when compared against itself. Got %d", s1)
	}
	if s2 < s1 {
		t.Error("Expected DL similarity to be greater when compared against itself")
	}
}

func TestLCP(t *testing.T) {
	assert := assert.New(t)
	lcp := LongestCommonPrefix("TELEKOMUNIKASI INDONESIA", "TELEKOMUNIKASI INDONESIA INTERNASIONAL")
	t.Log(lcp)
	assert.Equal("Hell", lcp)

	lcp = LongestCommonPrefix("Hello World", "Hell yeah!", "hey there")
	assert.Equal("", lcp)

	lcp = LongestCommonPrefix()
	assert.Equal("", lcp)

	lcp = LongestCommonPrefix("OneWord")
	t.Logf("LCP value : %v\n", lcp)
	assert.Equal("OneWord", lcp)

	lcp = LongestCommonPrefix("foo", "foobar")
	assert.Equal("foo", lcp)
}

var parseNumTests = []struct {
	s string
	v int
}{
	{"seratus dua puluh sembilan", 129},
	{"satu juta dua ratus lima puluh ribu", 1250000},
	{"dua ratus lima puluh ribu lima ratus lima puluh", 250550},
}

func TestParseNumber(t *testing.T) {
	for _, pnts := range parseNumTests {
		s := strings.Split(pnts.s, " ")
		ints, err := StrsToInts(s)
		if err != nil {
			t.Error(err)
			continue
		}

		v := CombineInts(ints)
		if v != pnts.v {
			t.Errorf("Expected %q to be parsed to %d. Got %d instead", pnts.s, pnts.v, v)
		}
	}
}

type Comparer struct {
	a string
	b string
}

func TestSimilarity(t *testing.T) {

	corpuses := []Comparer{
		{"TELEKOMUNIKASI INDONESIA", "TELEKOMUNIKASI INDONESIA INTERNASIONAL"},
		{"BASUKI PRATAMA ENGINEERING", "BASUKI PRATAMA ENGINERING"},
		{"BASUKI PRATAMA ENGINERING", "BASUKI PRATAMA ENGINNERING"},
		{"KONSORSIUM ANDHIKA ENERGINDO", "KONSORSIUM ANDIKA ENERGINDO"},
		{"MEGA ENERGI KHATULISTIWA", "MEGA ENERGY KHATULISTIWA"},
		{"MINAHASA CAHAYA LESTARI", "MITRA ENERGI BATAM"},
		{"MEDCO POWER INDONESIA", "MEDCO SARANA KALIBARU"},
	}

	for _, corpus := range corpuses {
		t.Log("Testing using DamerauLevenshtein.....")
		dl := DamerauLevenshtein(corpus.a, corpus.b)
		t.Log("Testing using Longest Common Prefix ...")
		lcp := LongestCommonPrefix(corpus.a, corpus.b)
		t.Log("Testing using CosineSimilarity ...")

		sliceA := strings.Split(corpus.a, " ")
		sliceB := strings.Split(corpus.b, " ")
		cs := CosineSimilarity(sliceA, sliceB)
		jw := JaroWinkler(corpus.a, corpus.b, true)
		jwf := JaroWinkler(corpus.a, corpus.b, false)
		t.Log(corpus.a, "| ", corpus.b)
		t.Logf("result comparision, DL method=%d, LCP method= %s, CS method=%f, JW method=%f, JW with false prefix=%f", dl, lcp, cs, jw, jwf)
	}

}
