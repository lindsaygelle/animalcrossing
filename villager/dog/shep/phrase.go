package shep

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "baa baa baa"
	phraseCanadianFrench       string = "pedigree"
	phraseDutch                string = "schapie"
	phraseFrench               string = "pedigree"
	phraseGerman               string = "heffheff"
	phraseItalian              string = "guau guau"
	phraseJapanese             string = "ワンダー"
	phraseLatinAmericanSpanish string = "grrruau"
	phraseKorean               string = "안보여"
	phraseRussian              string = "бр-р-разер"
	phraseSpanish              string = "tufos"
	phraseSimplifiedChinese    string = "汪想想"
	phraseTraditionalChinese   string = "汪想想"
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:      phraseAmericanEnglish,
		language.CanadianFrench:       phraseCanadianFrench,
		language.Dutch:                phraseDutch,
		language.French:               phraseFrench,
		language.German:               phraseGerman,
		language.Italian:              phraseItalian,
		language.Japanese:             phraseJapanese,
		language.Korean:               phraseKorean,
		language.LatinAmericanSpanish: phraseLatinAmericanSpanish,
		language.Russian:              phraseRussian,
		language.Spanish:              phraseSpanish,
		language.SimplifiedChinese:    phraseSimplifiedChinese,
		language.TraditionalChinese:   phraseTraditionalChinese}
)
