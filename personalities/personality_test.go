package personalities

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPersonalityName(t *testing.T) {
	var personalityName string = fmt.Sprint(rand.Int())
	var personality Personality = personality{
		name: personalityName}
	if ok := personality.Name() == personalityName; !ok {
		t.Fatalf("Personality.Name() != %s", personalityName)
	}
}
