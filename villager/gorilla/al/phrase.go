package al

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ayyyeee"
	phraseCanadianFrench       string = "ouba-oub"
	phraseDutch                string = "aiaiai"
	phraseFrench               string = "ouba-oub"
	phraseGerman               string = "uga-uga"
	phraseItalian              string = "bananao"
	phraseJapanese             string = "ウヒョッ"
	phraseLatinAmericanSpanish string = "juju-já"
	phraseKorean               string = "부락"
	phraseRussian              string = "ух-ух-ух"
	phraseSpanish              string = "chocolate"
	phraseSimplifiedChinese    string = "呜呼"
	phraseTraditionalChinese   string = "嗚呼"
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
