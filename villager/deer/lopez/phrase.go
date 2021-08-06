package lopez

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "badoom"
	phraseCanadianFrench       string = "you bet"
	phraseDutch                string = "bokje"
	phraseFrench               string = "jaunisse"
	phraseGerman               string = "bocko"
	phraseItalian              string = "sdleng"
	phraseJapanese             string = "がぜん"
	phraseLatinAmericanSpanish string = "trisca"
	phraseKorean               string = "엣헴"
	phraseRussian              string = "бубух"
	phraseSpanish              string = "trisca"
	phraseSimplifiedChinese    string = "变了"
	phraseTraditionalChinese   string = "變了"
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
