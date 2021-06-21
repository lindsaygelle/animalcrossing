package animals

import "testing"

func TestSeaGullName(t *testing.T) {
	if ok := SeaGull.Name() == seaGull; !ok {
		t.Fatal("SeaGull.Name() != seaGull")
	}
}
