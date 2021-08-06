package broccolo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "eat it"
	phraseCanadianFrench       string = "flim-flam"
	phraseDutch                string = "smakelijk"
	phraseFrench               string = "flim-flam"
	phraseGerman               string = "knabba"
	phraseItalian              string = "squikko"
	phraseJapanese             string = "ピコ"
	phraseLatinAmericanSpanish string = "flinflán"
	phraseKorean               string = "앗차"
	phraseRussian              string = "ням-ням"
	phraseSpanish              string = "dientes"
	phraseSimplifiedChinese    string = "微微"
	phraseTraditionalChinese   string = "微微"
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
