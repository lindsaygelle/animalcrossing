package rodeo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chaps"
	phraseCanadianFrench       string = "olé"
	phraseDutch                string = "cowboy"
	phraseFrench               string = "crac"
	phraseGerman               string = "oléolé"
	phraseItalian              string = "oléolé"
	phraseJapanese             string = "さすがに"
	phraseLatinAmericanSpanish string = "olé"
	phraseKorean               string = "디용"
	phraseRussian              string = "ковбой"
	phraseSpanish              string = "olé"
	phraseSimplifiedChinese    string = "了不起"
	phraseTraditionalChinese   string = "了不起"
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
