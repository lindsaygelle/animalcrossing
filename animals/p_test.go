package animals

import (
	"strings"
	"testing"
)

func TestP(t *testing.T) {
	for i := 0; i < len(P); i++ {
		var s string = strings.Title(P[i])
		if ok := P[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", P[i], s, i)
		}
		if ok := s[0] == 'P'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "P", i)
		}
	}
}
