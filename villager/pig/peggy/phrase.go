package peggy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "shweetie"
	phraseCanadianFrench       string = "fleur"
	phraseDutch                string = "knorrelief"
	phraseFrench               string = "fleur"
	phraseGerman               string = "glotz"
	phraseItalian              string = "oink"
	phraseJapanese             string = "ぷるる"
	phraseLatinAmericanSpanish string = "trufits"
	phraseKorean               string = "아앙"
	phraseRussian              string = "хрюмпатяга"
	phraseSpanish              string = "trufita"
	phraseSimplifiedChinese    string = "弹弹"
	phraseTraditionalChinese   string = "彈彈"
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
