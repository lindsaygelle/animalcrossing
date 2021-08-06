package maggie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "schep"
	phraseCanadianFrench       string = "grouik"
	phraseDutch                string = "knor-knor"
	phraseFrench               string = "grouik"
	phraseGerman               string = "kringel"
	phraseItalian              string = "sissì"
	phraseJapanese             string = "うん"
	phraseLatinAmericanSpanish string = "oinki"
	phraseKorean               string = "유후"
	phraseRussian              string = "хрюля"
	phraseSpanish              string = "girasol"
	phraseSimplifiedChinese    string = "嗯"
	phraseTraditionalChinese   string = "嗯"
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
