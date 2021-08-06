package ozzie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ol' bear"
	phraseCanadianFrench       string = "koa là là"
	phraseDutch                string = "ouwe beer"
	phraseFrench               string = "koa là là"
	phraseGerman               string = "altes Haus"
	phraseItalian              string = "ulabadula"
	phraseJapanese             string = "ククッ"
	phraseLatinAmericanSpanish string = "koahhh"
	phraseKorean               string = "큐큐"
	phraseRussian              string = "медведище"
	phraseSpanish              string = "osete"
	phraseSimplifiedChinese    string = "颗颗"
	phraseTraditionalChinese   string = "顆顆"
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
