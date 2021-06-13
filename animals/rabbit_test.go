package animals

import "testing"

func TestRabbitName(t *testing.T) {
	if ok := Rabbit.Name() == rabbit; !ok {
		t.Fatal("Rabbit.Name() != rabbit")
	}
}
