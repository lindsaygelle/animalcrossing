package boomer

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "human"
	phraseCanadianFrench       string = "humanoïde"
	phraseDutch                string = "mens"
	phraseFrench               string = "boudiou"
	phraseGerman               string = "menschlein"
	phraseItalian              string = "glacius"
	phraseJapanese             string = "ツーツツ"
	phraseLatinAmericanSpanish string = "glaciux"
	phraseKorean               string = "에헴헴"
	phraseRussian              string = "человек"
	phraseSpanish              string = "humanoide"
	phraseSimplifiedChinese    string = "轧轧"
	phraseTraditionalChinese   string = "軋軋"
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
