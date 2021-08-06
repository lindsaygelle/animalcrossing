package butch

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ROOOOOWF"
	phraseCanadianFrench       string = "WROOOOUF"
	phraseDutch                string = "BROEEEF"
	phraseFrench               string = "WROOOOUF"
	phraseGerman               string = "GRRUFF"
	phraseItalian              string = "ROOOOOF"
	phraseJapanese             string = "ノン"
	phraseLatinAmericanSpanish string = "gruf-gruf"
	phraseKorean               string = "아니"
	phraseRussian              string = "ррр-гав"
	phraseSpanish              string = "gruf-gruf"
	phraseSimplifiedChinese    string = "侬"
	phraseTraditionalChinese   string = "儂"
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
