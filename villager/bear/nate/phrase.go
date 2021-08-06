package nate

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "yawwwn"
	phraseCanadianFrench       string = "baaîîîllle"
	phraseDutch                string = "gaaaap"
	phraseFrench               string = "baaîîîllle"
	phraseGerman               string = "gäähn"
	phraseItalian              string = "yauuun"
	phraseJapanese             string = "んだ"
	phraseLatinAmericanSpanish string = "uuaaaah"
	phraseKorean               string = "힘내"
	phraseRussian              string = "спа-а-ать"
	phraseSpanish              string = "uuaaaah"
	phraseSimplifiedChinese    string = "嗯是"
	phraseTraditionalChinese   string = "嗯是"
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
