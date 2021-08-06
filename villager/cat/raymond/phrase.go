package raymond

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "crisp"
	phraseCanadianFrench       string = "panache"
	phraseDutch                string = "netjes"
	phraseFrench               string = "panache"
	phraseGerman               string = "weeßte"
	phraseItalian              string = "rrricooo"
	phraseJapanese             string = "キリッ"
	phraseLatinAmericanSpanish string = "atilda"
	phraseKorean               string = "우쭐"
	phraseRussian              string = "стильненько"
	phraseSpanish              string = "atilda"
	phraseSimplifiedChinese    string = "严肃"
	phraseTraditionalChinese   string = "嚴肅"
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
