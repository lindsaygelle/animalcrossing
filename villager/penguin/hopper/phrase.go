package hopper

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "slushie"
	phraseCanadianFrench       string = "manette"
	phraseDutch                string = "ĳsje"
	phraseFrench               string = "manette"
	phraseGerman               string = "slaschi"
	phraseItalian              string = "sguish"
	phraseJapanese             string = "ちぅねん"
	phraseLatinAmericanSpanish string = "esploch"
	phraseKorean               string = "아따"
	phraseRussian              string = "ледок"
	phraseSpanish              string = "esploch"
	phraseSimplifiedChinese    string = "禅"
	phraseTraditionalChinese   string = "禪"
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
