package animals

import "testing"

func TestCamelName(t *testing.T) {
	if ok := Camel.Name() == camel; !ok {
		t.Fatal("Camel.Name() != camel")
	}
}
