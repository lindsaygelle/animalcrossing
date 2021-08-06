package holden

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "glue stick"
	phraseCanadianFrench       string = "gluglu"
	phraseDutch                string = ""
	phraseFrench               string = "gluglu"
	phraseGerman               string = "klebstift"
	phraseItalian              string = "incollo"
	phraseJapanese             string = "ノリノリ"
	phraseLatinAmericanSpanish string = "tudiquesí"
	phraseKorean               string = "둠칫둠칫"
	phraseRussian              string = ""
	phraseSpanish              string = "tudiquesí"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
