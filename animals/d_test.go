package animals

import (
	"strings"
	"testing"
)

func TestD(t *testing.T) {
	for i := 0; i < len(D); i++ {
		var s string = strings.Title(D[i])
		if ok := D[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", D[i], s, i)
		}
		if ok := s[0] == 'D'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "D", i)
		}
	}
}