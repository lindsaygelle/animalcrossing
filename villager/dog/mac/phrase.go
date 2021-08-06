package mac

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "woo woof"
	phraseCanadianFrench       string = "babines"
	phraseDutch                string = "woewoef"
	phraseFrench               string = "babines"
	phraseGerman               string = "bellbell"
	phraseItalian              string = "arfidenti"
	phraseJapanese             string = "モグ"
	phraseLatinAmericanSpanish string = "frusky"
	phraseKorean               string = "우물우물"
	phraseRussian              string = "ау-у-гав"
	phraseSpanish              string = "frusky"
	phraseSimplifiedChinese    string = "摘"
	phraseTraditionalChinese   string = "摘"
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
