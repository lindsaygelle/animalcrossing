package peck

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "crunch"
	phraseCanadianFrench       string = "poupousse"
	phraseDutch                string = "wasbordje"
	phraseFrench               string = "poupousse"
	phraseGerman               string = "flatter"
	phraseItalian              string = "sprint"
	phraseJapanese             string = "ブン"
	phraseLatinAmericanSpanish string = "cachirp"
	phraseKorean               string = "츠츠츠"
	phraseRussian              string = "тук-тук"
	phraseSpanish              string = "muy fuerte"
	phraseSimplifiedChinese    string = "噗噗"
	phraseTraditionalChinese   string = "噗噗"
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
