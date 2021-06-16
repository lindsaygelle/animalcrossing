package astrology

import "testing"

func TestLeoName(t *testing.T) {
	if ok := Leo.Name() == leo; !ok {
		t.Fatalf("%s != %s", Leo.Name(), leo)
	}
}
