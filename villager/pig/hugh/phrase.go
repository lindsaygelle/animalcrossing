package hugh

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snortle"
	phraseCanadianFrench       string = "cuick"
	phraseDutch                string = "snuif"
	phraseFrench               string = "cuick"
	phraseGerman               string = "schniff"
	phraseItalian              string = "snort"
	phraseJapanese             string = "とかね"
	phraseLatinAmericanSpanish string = "groink"
	phraseKorean               string = "아님말구"
	phraseRussian              string = "хрюк-хрюк"
	phraseSpanish              string = "groink"
	phraseSimplifiedChinese    string = "懒懒"
	phraseTraditionalChinese   string = "懶懶"
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
