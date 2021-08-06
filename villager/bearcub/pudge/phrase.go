package pudge

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "golly"
	phraseCanadianFrench       string = "louloute"
	phraseDutch                string = "pudding"
	phraseFrench               string = "louloute"
	phraseGerman               string = "griesgram"
	phraseItalian              string = "pulcino"
	phraseJapanese             string = "んもう"
	phraseLatinAmericanSpanish string = "bruah"
	phraseKorean               string = "아이참"
	phraseRussian              string = "карапуз"
	phraseSpanish              string = "rebollos"
	phraseSimplifiedChinese    string = "真是的"
	phraseTraditionalChinese   string = "真是的"
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
