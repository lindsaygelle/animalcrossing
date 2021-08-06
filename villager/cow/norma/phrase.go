package norma

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hoof hoo"
	phraseCanadianFrench       string = "meuh nan"
	phraseDutch                string = "boehoe"
	phraseFrench               string = "meuh nan"
	phraseGerman               string = "muuuhi"
	phraseItalian              string = "mu muuu"
	phraseJapanese             string = "うふ"
	phraseLatinAmericanSpanish string = "mu-mu"
	phraseKorean               string = "에헤"
	phraseRussian              string = "до-ре-му"
	phraseSpanish              string = "mu-mu"
	phraseSimplifiedChinese    string = "微笑"
	phraseTraditionalChinese   string = "微笑"
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
