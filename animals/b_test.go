package animals

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	for i := 0; i < len(B); i++ {
		var s string = strings.Title(B[i])
		if ok := B[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", B[i], s, i)
		}
		if ok := s[0] == 'B'; !ok {
			t.Fatalf("%s != %s; i=%d", string(s[0]), "B", i)
		}
	}
}