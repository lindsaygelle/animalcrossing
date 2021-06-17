package animals

import (
	"strings"
	"testing"
)

func TestT(t *testing.T) {
	for i := 0; i < len(T); i++ {
		var s string = strings.Title(T[i])
		if ok := T[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", T[i], s, i)
		}
	}
}
