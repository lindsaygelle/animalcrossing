package animals

import "testing"

func TestLionName(t *testing.T) {
	if ok := Lion.Name() == lion; !ok {
		t.Fatal("Lion.Name() != lion")
	}
}
