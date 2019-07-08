package corpus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
	assert := assert.New(t)
	dict := New()
	assert.Equal(0, dict.WordFreq("selamat")) // frequency of a word not in dict ould have to be 0
	assert.Equal(0, dict.IDFreq(3))           // ditto

	id := dict.Add("selamat")

	assert.Equal(3, id)
	assert.Equal([]string{"", "-UNKNOWN-", "-ROOT-", "selamat"}, dict.words)
	assert.Equal(map[string]int{"": 0, "-UNKNOWN-": 1, "-ROOT-": 2, "selamat": 3}, dict.ids)
	assert.Equal(4, dict.Size())

	id2, ok := dict.Id("selamat")
	if !ok {
		t.Errorf("The ID of null should be  0")
	}
	assert.Equal(id, id2)

	word, ok := dict.Word(3)
	if !ok {
		t.Errorf("Expected word of ID 3 to be found")
	}
	assert.Equal("selamat", word)

	dict.Add(word)
	assert.Equal(2, dict.WordFreq(word))
	assert.Equal(2, dict.IDFreq(3))
	assert.Equal(5, dict.TotalFreq())
	assert.Equal(7, dict.MaxWordLength())

	prob, ok := dict.WordProb(word)
	if !ok {
		t.Errorf("Expected a probability")
	}
	assert.Equal(0.4, prob)
	// t.Logf("%q: %v", word, prob)
}

func TestCorpus_Merge(t *testing.T) {
	assert := assert.New(t)

	dict := New()
	id := dict.Add("selamat")
	dict.frequencies[id] += 4 // freq for "hello" is 5
	dict.totalFreq += 4

	other := New()
	id = other.Add("selamat")
	other.frequencies[id] += 2 // freq for "hello" is 3
	other.totalFreq += 2
	id = other.Add("malam")
	other.frequencies[id] += 1
	other.totalFreq += 1

	dict.Merge(other)

	assert.Equal(8, dict.WordFreq("selamat"))
	assert.Equal(2, dict.WordFreq("malam"))
}
