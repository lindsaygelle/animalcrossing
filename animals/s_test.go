package animals

import (
	"strings"
	"testing"
)

func TestS(t *testing.T) {
	for i := 0; i < len(S); i++ {
		var s string = strings.Title(S[i])
		if ok := S[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", S[i], s, i)
		}
	}
}
