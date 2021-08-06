package vladimir

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nyet"
	phraseCanadianFrench       string = "niet"
	phraseDutch                string = "njet"
	phraseFrench               string = "niet"
	phraseGerman               string = "nastrovje"
	phraseItalian              string = "nyet"
	phraseJapanese             string = "やんけ"
	phraseLatinAmericanSpanish string = "nyet"
	phraseKorean               string = "참말로"
	phraseRussian              string = "ну не-е-ет"
	phraseSpanish              string = "nyet"
	phraseSimplifiedChinese    string = "唉唷"
	phraseTraditionalChinese   string = "唉唷"
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
