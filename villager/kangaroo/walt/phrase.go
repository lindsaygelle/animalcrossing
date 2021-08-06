package walt

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pockets"
	phraseCanadianFrench       string = "popoche"
	phraseDutch                string = "zakkies"
	phraseFrench               string = "popoche"
	phraseGerman               string = "boingboing"
	phraseItalian              string = "boing"
	phraseJapanese             string = "じゃのう"
	phraseLatinAmericanSpanish string = "boing"
	phraseKorean               string = "어험"
	phraseRussian              string = "карманчик"
	phraseSpanish              string = "boing"
	phraseSimplifiedChinese    string = "别装了"
	phraseTraditionalChinese   string = "別裝了"
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
