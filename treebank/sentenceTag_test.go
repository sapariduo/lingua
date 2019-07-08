package treebank

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentenceTag(t *testing.T) {
	assert := assert.New(t)
	readr := strings.NewReader(sampleConllu)
	st := ReadConllu(readr)[0]
	t.Logf("%+v", st.Labels)
	correctHeads := []int{2, 0, 2, 2, 4, 4, 9, 9, 2, 11, 9, 11, 2}
	assert.Equal(correctHeads, st.Heads)

	dep := st.Dependency(nil)
	t.Logf("%v", dep)

	assert.Equal(correctHeads, dep.Heads()[1:])
}
