package lily

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "toady"
	phraseCanadianFrench       string = "zzrrbbitt"
	phraseDutch                string = "padje"
	phraseFrench               string = "watich'"
	phraseGerman               string = "kröti"
	phraseItalian              string = "gragragra"
	phraseJapanese             string = "でちゅ"
	phraseLatinAmericanSpanish string = "riiibit"
	phraseKorean               string = "그래요"
	phraseRussian              string = "ква"
	phraseSpanish              string = "mosquita"
	phraseSimplifiedChinese    string = "雨雨"
	phraseTraditionalChinese   string = "雨雨"
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
