package animals

import "testing"

func TestReindeerName(t *testing.T) {
	if ok := Reindeer.Name() == reindeer; !ok {
		t.Fatal("Reindeer.Name() != reindeer")
	}
}
