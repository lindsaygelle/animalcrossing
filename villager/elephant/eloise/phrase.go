package eloise

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tooooot"
	phraseCanadianFrench       string = "tooouut"
	phraseDutch                string = "pwèèèèèp"
	phraseFrench               string = "tooouut"
	phraseGerman               string = "tuutuuuu"
	phraseItalian              string = "nappole"
	phraseJapanese             string = "ルン"
	phraseLatinAmericanSpanish string = "fruuuf"
	phraseKorean               string = "훨훨"
	phraseRussian              string = "тру-тут"
	phraseSpanish              string = "orejillas"
	phraseSimplifiedChinese    string = "鲁鲁"
	phraseTraditionalChinese   string = "魯魯"
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
