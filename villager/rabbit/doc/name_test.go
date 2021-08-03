package doc

import (
	"testing"

	"golang.org/x/text/language"
)

func TestName(t *testing.T) {
	tags := []language.Tag{
		language.AmericanEnglish,
		language.French,
		language.German,
		language.Italian,
		language.Japanese,
		language.Korean,
		language.Russian,
		language.Spanish,
		language.SimplifiedChinese,
		language.TraditionalChinese}
	for _, tag := range tags {
		name, ok := name[tag]
		s := tag.String()
		if !ok {
			t.Fatalf("name[%s] != true", s)
		}
		if ok := len(name) > 0; !ok {
			t.Skipf("len(name[%s]) == 0", s)
		}
	}
}