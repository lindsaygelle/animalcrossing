package hamphrey

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snort"
	phraseCanadianFrench       string = "gna gna"
	phraseDutch                string = "snuifsnuif"
	phraseFrench               string = "balivernes"
	phraseGerman               string = "squiep"
	phraseItalian              string = "spatapumf"
	phraseJapanese             string = "カーッ"
	phraseLatinAmericanSpanish string = "noquenó"
	phraseKorean               string = "카악"
	phraseRussian              string = "фырк"
	phraseSpanish              string = "andaquenó"
	phraseSimplifiedChinese    string = "怒"
	phraseTraditionalChinese   string = "怒"
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
