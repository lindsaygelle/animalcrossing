package animals

import (
	"strings"
	"testing"
)

func TestR(t *testing.T) {
	for i := 0; i < len(R); i++ {
		var s string = strings.Title(R[i])
		if ok := R[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", R[i], s, i)
		}
	}
}
