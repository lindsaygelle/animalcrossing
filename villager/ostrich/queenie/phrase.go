package queenie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chicken"
	phraseCanadianFrench       string = "poupoule"
	phraseDutch                string = "kippie"
	phraseFrench               string = "poupoule"
	phraseGerman               string = "hühnchen"
	phraseItalian              string = "polpetta"
	phraseJapanese             string = "やっぱし"
	phraseLatinAmericanSpanish string = "plumifá"
	phraseKorean               string = "역시"
	phraseRussian              string = "цыпленок"
	phraseSpanish              string = "gallina"
	phraseSimplifiedChinese    string = "果然"
	phraseTraditionalChinese   string = "果然"
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
