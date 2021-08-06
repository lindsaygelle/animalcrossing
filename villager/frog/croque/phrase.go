package croque

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "as if"
	phraseCanadianFrench       string = "ouaoua"
	phraseDutch                string = "ja dááág"
	phraseFrench               string = "crapoupou"
	phraseGerman               string = "schleck"
	phraseItalian              string = "umpf"
	phraseJapanese             string = "しからば"
	phraseLatinAmericanSpanish string = "yastá"
	phraseKorean               string = "툴툴"
	phraseRussian              string = "вроде того"
	phraseSpanish              string = "yastá"
	phraseSimplifiedChinese    string = "子曰"
	phraseTraditionalChinese   string = "子曰"
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
