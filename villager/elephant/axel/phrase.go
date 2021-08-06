package axel

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "WHONK"
	phraseCanadianFrench       string = "splaf"
	phraseDutch                string = "PWAAP"
	phraseFrench               string = "splaf"
	phraseGerman               string = "TUUUUT"
	phraseItalian              string = "SBONK"
	phraseJapanese             string = "でゴンス"
	phraseLatinAmericanSpanish string = "ankagua"
	phraseKorean               string = "히힛"
	phraseRussian              string = "потрубим"
	phraseSpanish              string = "ankagua"
	phraseSimplifiedChinese    string = "嘻嘻"
	phraseTraditionalChinese   string = "嘻嘻"
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
