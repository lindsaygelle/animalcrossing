package megumi

import (
	"testing"

	"golang.org/x/text/language"
)

func TestPhrase(t *testing.T) {
	tags := []language.Tag{
		language.AmericanEnglish,
		language.CanadianFrench,
		language.Dutch,
		language.French,
		language.German,
		language.Italian,
		language.Japanese,
		language.Korean,
		language.LatinAmericanSpanish,
		language.Russian,
		language.Spanish,
		language.SimplifiedChinese,
		language.TraditionalChinese}
	for _, tag := range tags {
		phrase, ok := phrase[tag]
		s := tag.String()
		if !ok {
			t.Fatalf("phrase[%s] != true", s)
		}
		if ok := len(phrase) > 0; !ok {
			t.Skipf("len(phrase[%s]) == 0", s)
		}
	}
}
