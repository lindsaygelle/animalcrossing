package animals

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	for i := 0; i < len(A); i++ {
		var s string = strings.Title(A[i])
		if ok := A[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", A[i], s, i)
		}
		if ok := s[0] == 'A'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "A", i)
		}
	}
}
