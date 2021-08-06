package anchovy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chuurp"
	phraseCanadianFrench       string = "muf-muf"
	phraseDutch                string = "tsjiiirp"
	phraseFrench               string = "muf-muf"
	phraseGerman               string = "tschurp"
	phraseItalian              string = "ciripì"
	phraseJapanese             string = "でシ"
	phraseLatinAmericanSpanish string = "piu-piu"
	phraseKorean               string = "이어라"
	phraseRussian              string = "чик-чирик"
	phraseSpanish              string = "piu-piu"
	phraseSimplifiedChinese    string = "如斯"
	phraseTraditionalChinese   string = "如斯"
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
