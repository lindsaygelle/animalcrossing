package joey

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bleeeeeck"
	phraseCanadianFrench       string = "beeeuuuurk"
	phraseDutch                string = "kèèèk"
	phraseFrench               string = "beeeuuuurk"
	phraseGerman               string = "tschiiirp"
	phraseItalian              string = "squack"
	phraseJapanese             string = "でヤンス"
	phraseLatinAmericanSpanish string = "uac-uac"
	phraseKorean               string = "그래유"
	phraseRussian              string = "кряк"
	phraseSpanish              string = "uac-uac"
	phraseSimplifiedChinese    string = "鸭鸭"
	phraseTraditionalChinese   string = "鴨鴨"
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
