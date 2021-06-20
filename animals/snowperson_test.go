package animals

import "testing"

func TestSnowpersonName(t *testing.T) {
	if ok := Snowperson.Name() == snowperson; !ok {
		t.Fatal("Snowperson.Name() != snowperson")
	}
}
