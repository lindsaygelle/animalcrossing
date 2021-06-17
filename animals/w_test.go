package animals

import (
	"strings"
	"testing"
)

func TestW(t *testing.T) {
	for i := 0; i < len(W); i++ {
		var s string = strings.Title(W[i])
		if ok := W[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", W[i], s, i)
		}
		if ok := s[0] == 'W'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "W", i)
		}
	}
}
