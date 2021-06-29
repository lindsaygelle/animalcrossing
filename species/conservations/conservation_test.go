package conservations

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	var (
		c = [...]string{
			criticallyEndangered,
			domesticated,
			endangered,
			extinct,
			leastConcern,
			nearThreatened,
			unknown,
			vulnerable}
	)

	for i := 0; i < len(c); i++ {
		var s string = c[i]
		if ok := string(s[0]) == strings.ToUpper(string(s[0])); !ok {
			t.Fatalf("%s does not start with a capital letter", s)
		}
	}
}
