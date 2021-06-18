package species

import (
	"strings"
	"testing"
)

func TestM(t *testing.T) {
	for i := 0; i < len(M); i++ {
		var s string = strings.Title(M[i])
		if ok := M[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", M[i], s, i)
		}
		if ok := s[0] == 'M'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "M", i)
		}
	}
}
