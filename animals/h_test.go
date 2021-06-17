package animals

import (
	"strings"
	"testing"
)

func TestH(t *testing.T) {
	for i := 0; i < len(H); i++ {
		var s string = strings.Title(H[i])
		if ok := H[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", H[i], s, i)
		}
	}
}
