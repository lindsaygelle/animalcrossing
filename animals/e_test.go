package animals

import (
	"strings"
	"testing"
)

func TestE(t *testing.T) {
	for i := 0; i < len(E); i++ {
		var s string = strings.Title(E[i])
		if ok := E[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", E[i], s, i)
		}
	}
}
