package sylvia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "boing"
	phraseCanadianFrench       string = "d'la balle"
	phraseDutch                string = "boing"
	phraseFrench               string = "d'la balle"
	phraseGerman               string = "flottflott"
	phraseItalian              string = "saltolà"
	phraseJapanese             string = "よねー"
	phraseLatinAmericanSpanish string = "do-doing"
	phraseKorean               string = "이자슥"
	phraseRussian              string = "прыг"
	phraseSpanish              string = "púgil"
	phraseSimplifiedChinese    string = "对耶"
	phraseTraditionalChinese   string = "對耶"
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
