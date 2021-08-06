package sterling

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "skraaaaw"
	phraseCanadianFrench       string = "kaputt"
	phraseDutch                string = "havik"
	phraseFrench               string = "kaputt"
	phraseGerman               string = "krahkrah"
	phraseItalian              string = "chicchirò"
	phraseJapanese             string = "やあッ"
	phraseLatinAmericanSpanish string = "escróóó"
	phraseKorean               string = "호옷"
	phraseRussian              string = "клекот"
	phraseSpanish              string = "mancuernas"
	phraseSimplifiedChinese    string = "呀啊"
	phraseTraditionalChinese   string = "呀啊"
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
