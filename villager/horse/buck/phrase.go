package buck

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pardner"
	phraseCanadianFrench       string = "howdy"
	phraseDutch                string = "draver"
	phraseFrench               string = "taty"
	phraseGerman               string = "wiehie"
	phraseItalian              string = "cowboy"
	phraseJapanese             string = "なんちて"
	phraseLatinAmericanSpanish string = "jopop"
	phraseKorean               string = "농담"
	phraseRussian              string = "партнер"
	phraseSpanish              string = "zanahoria"
	phraseSimplifiedChinese    string = "搞笑"
	phraseTraditionalChinese   string = "搞笑"
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
