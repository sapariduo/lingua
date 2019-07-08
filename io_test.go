package lingua

import (
	"encoding/json"
	"testing"
)

func TestAnnotationJSON(t *testing.T) {
	a := NewAnnotation()
	a.Value = "sudah"
	a.POSTag = ADV
	a.DependencyType = AdvMod
	a.ID = 2

	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	t.Logf(" %s", string(b))

	x := `{"ID":2,"Value":"sudah","POSTag":"ADV","Label":"AdvMod"}`
	c := NewAnnotation()
	if err = json.Unmarshal([]byte(x), c); err != nil {
		t.Error(err)
	}

	if c.Value != a.Value {
		t.Errorf("Expected Value to be %q. Got %q insteed", a.Value, c.Value)
	}

	if c.POSTag != a.POSTag {
		t.Errorf("Expected POSTag to be %v. Got %v instead", a.POSTag, c.POSTag)
	}

	if c.DependencyType != a.DependencyType {
		t.Errorf("Expected DependencyType to be %v. Got %v instead", a.DependencyType, c.DependencyType)
	}
}

func TestAnnotatedSentenceJSON(t *testing.T) {
	a := NewAnnotation()
	a.Value = "sudah"
	a.POSTag = ADV
	a.DependencyType = AdvMod
	a.ID = 0

	b := NewAnnotation()
	b.Value = "lama"
	b.POSTag = ADJ
	b.DependencyType = AMod
	b.ID = 1
	b.Head = rootAnnotation

	a.Head = b

	as := AnnotatedSentence{a, b}
	bs, err := json.Marshal(as)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(bs))

	x := `[{"ID":0,"Value":"sudah","POSTag":"ADV","Label":"AdvMod","Head":1},{"ID":1,"Value":"lama","POSTag":"ADJ","Label":"AMod","Head":-1000}]`

	var cs AnnotatedSentence
	if err = json.Unmarshal([]byte(x), &cs); err != nil {
		t.Error(err)
	}
	t.Logf("%v", cs)

	for i, c := range cs {
		d := as[i]

		if c.Value != d.Value {
			t.Error("Expected Values to be the same")
		}

		if c.POSTag != d.POSTag {
			t.Error("POSTag not the same")
		}

		if c.DependencyType != d.DependencyType {
			t.Error("Dependency Types not the same")
		}

		if c.HeadID() != d.HeadID() {
			t.Errorf("%v HeadIDs not the same. Want %v, got %v instead", d, d.HeadID(), c.HeadID())
		}
	}
}
