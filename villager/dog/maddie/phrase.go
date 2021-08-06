package maddie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "yippee"
	phraseCanadianFrench       string = "youpi"
	phraseDutch                string = "jippie"
	phraseFrench               string = "youp"
	phraseGerman               string = "hurra"
	phraseItalian              string = "yuuuppii"
	phraseJapanese             string = "ワンツー"
	phraseLatinAmericanSpanish string = "reguau"
	phraseKorean               string = "달코미"
	phraseRussian              string = "ур-ра"
	phraseSpanish              string = "chupi"
	phraseSimplifiedChinese    string = "一二"
	phraseTraditionalChinese   string = "一二"
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
