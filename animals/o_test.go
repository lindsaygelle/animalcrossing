package animals

import (
	"strings"
	"testing"
)

func TestO(t *testing.T) {
	for i := 0; i < len(O); i++ {
		var s string = strings.Title(O[i])
		if ok := O[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", O[i], s, i)
		}
	}
}
