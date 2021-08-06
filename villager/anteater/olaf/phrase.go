package olaf

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "whiffa"
	phraseCanadianFrench       string = "yesss"
	phraseDutch                string = "toet"
	phraseFrench               string = "tartarin"
	phraseGerman               string = "mwahaaa"
	phraseItalian              string = "swish"
	phraseJapanese             string = "ムーチョ"
	phraseLatinAmericanSpanish string = "fruf"
	phraseKorean               string = "무쵸"
	phraseRussian              string = "уф-уф"
	phraseSpanish              string = "fruf"
	phraseSimplifiedChinese    string = "欧啦"
	phraseTraditionalChinese   string = "歐啦"
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
