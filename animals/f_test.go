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
	}
}
