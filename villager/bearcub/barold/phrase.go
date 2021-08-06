package barold

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cubby"
	phraseCanadianFrench       string = "dodododu"
	phraseDutch                string = "welpje"
	phraseFrench               string = "dodododu"
	phraseGerman               string = "mwuuu"
	phraseItalian              string = "babalù"
	phraseJapanese             string = "いっそ"
	phraseLatinAmericanSpanish string = "apurri"
	phraseKorean               string = "아니그냥"
	phraseRussian              string = "берлога"
	phraseSpanish              string = "apurri"
	phraseSimplifiedChinese    string = "算了"
	phraseTraditionalChinese   string = "算了"
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
