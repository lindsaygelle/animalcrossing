package animals

import (
	"strings"
	"testing"
)

func TestF(t *testing.T) {
	for i := 0; i < len(F); i++ {
		var s string = strings.Title(F[i])
		if ok := F[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", F[i], s, i)
		}
		if ok := s[0] == 'F'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "F", i)
		}
	}
}
