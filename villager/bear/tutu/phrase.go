package tutu

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "twinkles"
	phraseCanadianFrench       string = "pic pic"
	phraseDutch                string = "honingbij"
	phraseFrench               string = "pic pic"
	phraseGerman               string = "hooonig"
	phraseItalian              string = "gnifa"
	phraseJapanese             string = "ファイト"
	phraseLatinAmericanSpanish string = "fisflash"
	phraseKorean               string = "사르르"
	phraseRussian              string = "ме-е-ед"
	phraseSpanish              string = "fisflash"
	phraseSimplifiedChinese    string = "加油"
	phraseTraditionalChinese   string = "加油"
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
