package wendy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "lambkins"
	phraseCanadianFrench       string = "bê-bê-bê"
	phraseDutch                string = "lammetje"
	phraseFrench               string = "bê-bê-bê"
	phraseGerman               string = "wuschel"
	phraseItalian              string = "beee-ehi"
	phraseJapanese             string = "イシシ"
	phraseLatinAmericanSpanish string = "bee-biis"
	phraseKorean               string = "꺄르르"
	phraseRussian              string = "овечка"
	phraseSpanish              string = "joyita"
	phraseSimplifiedChinese    string = "一丝丝"
	phraseTraditionalChinese   string = "一絲絲"
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
