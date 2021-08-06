package raddle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "aaach—"
	phraseCanadianFrench       string = "kroah"
	phraseDutch                string = "kraah"
	phraseFrench               string = "kroah"
	phraseGerman               string = "tatü"
	phraseItalian              string = "rrruuuak"
	phraseJapanese             string = "へくしっ"
	phraseLatinAmericanSpanish string = "croaqui"
	phraseKorean               string = "콜록"
	phraseRussian              string = "ах-ох"
	phraseSpanish              string = "croaqui"
	phraseSimplifiedChinese    string = "哈啾"
	phraseTraditionalChinese   string = "哈啾"
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
