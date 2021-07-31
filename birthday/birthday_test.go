package birthday

import (
	"math/rand"
	"testing"
	"time"
)

// TestBirthdayDay tests whether the Birthday has the correct day.
func TestBirthdayDay(t *testing.T) {
	if v := NewBirthday(1, time.January); !(v.Day() == 1) {
		t.Fatalf("%d != %d", v.Day(), 1)
	}
}

// TestBirthdayMonth tests whether the Birthday has the correct month.
func TestBirthdayMonth(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	month := time.Month(rand.Intn(12-1+1) + 1)
	if v := NewBirthday(1, month); !(v.Month() == month) {
		t.Fatalf("%s != %s", v.Month(), month)
	}
}
