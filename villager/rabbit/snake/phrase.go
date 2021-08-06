package snake

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bunyip"
	phraseCanadianFrench       string = "mon lapin"
	phraseDutch                string = "ninja"
	phraseFrench               string = "mon lapin"
	phraseGerman               string = "hopplahopp"
	phraseItalian              string = "bun bun"
	phraseJapanese             string = "ニンニン"
	phraseLatinAmericanSpanish string = "boin-boing"
	phraseKorean               string = "닌닌"
	phraseRussian              string = "заяссс"
	phraseSpanish              string = "boin-boing"
	phraseSimplifiedChinese    string = "忍忍"
	phraseTraditionalChinese   string = "忍忍"
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
