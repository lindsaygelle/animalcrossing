package yuka

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tsk tsk"
	phraseCanadianFrench       string = "tss tss"
	phraseDutch                string = "tss tss"
	phraseFrench               string = "tss tss"
	phraseGerman               string = "grins"
	phraseItalian              string = "tsé tsé"
	phraseJapanese             string = "アラ"
	phraseLatinAmericanSpanish string = "tsé tsé"
	phraseKorean               string = "어맛"
	phraseRussian              string = "ц-ц"
	phraseSpanish              string = "pues"
	phraseSimplifiedChinese    string = "啊啦"
	phraseTraditionalChinese   string = "啊啦"
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
