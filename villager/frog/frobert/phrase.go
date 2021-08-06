package frobert

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "fribbit"
	phraseCanadianFrench       string = "breup"
	phraseDutch                string = "fffwaak"
	phraseFrench               string = "breup"
	phraseGerman               string = "quäääk"
	phraseItalian              string = "fribbit"
	phraseJapanese             string = "クルリ"
	phraseLatinAmericanSpanish string = "fribit"
	phraseKorean               string = "객울"
	phraseRussian              string = "квак"
	phraseSpanish              string = "anquitas"
	phraseSimplifiedChinese    string = "咕噜"
	phraseTraditionalChinese   string = "咕嚕"
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
