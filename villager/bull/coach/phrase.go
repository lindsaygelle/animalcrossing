package coach

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "stubble"
	phraseCanadianFrench       string = "bétail"
	phraseDutch                string = "stoppel"
	phraseFrench               string = "mon cou"
	phraseGerman               string = "schnoff"
	phraseItalian              string = "oooh issa"
	phraseJapanese             string = "ジョリッ"
	phraseLatinAmericanSpanish string = "chacho"
	phraseKorean               string = "땅땅"
	phraseRussian              string = "раз-два"
	phraseSpanish              string = "chacho"
	phraseSimplifiedChinese    string = "胡渣"
	phraseTraditionalChinese   string = "鬍渣"
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
