package animals

import (
	"strings"
	"testing"
)

func TestG(t *testing.T) {
	for i := 0; i < len(G); i++ {
		var s string = strings.Title(G[i])
		if ok := G[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", G[i], s, i)
		}
		if ok := s[0] == 'G'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "G", i)
		}
	}
}
