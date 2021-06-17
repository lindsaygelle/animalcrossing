package animals

import (
	"strings"
	"testing"
)

func TestK(t *testing.T) {
	for i := 0; i < len(K); i++ {
		var s string = strings.Title(K[i])
		if ok := K[i] == s; !ok {
			t.Fatalf("%s != %s; i=%d", K[i], s, i)
		}
	}
}
