package skye

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "airmail"
	phraseCanadianFrench       string = "mumuseau"
	phraseDutch                string = "luchtpost"
	phraseFrench               string = "mumuseau"
	phraseGerman               string = "ahiii"
	phraseItalian              string = "uuuhlalà"
	phraseJapanese             string = "フワワ"
	phraseLatinAmericanSpanish string = "auuu"
	phraseKorean               string = "후와와"
	phraseRussian              string = "пилот"
	phraseSpanish              string = "auuu"
	phraseSimplifiedChinese    string = "花花"
	phraseTraditionalChinese   string = "花花"
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
