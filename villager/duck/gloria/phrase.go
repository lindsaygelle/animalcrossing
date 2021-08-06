package gloria

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quacker"
	phraseCanadianFrench       string = "kwakos"
	phraseDutch                string = "kwebbel"
	phraseFrench               string = "kwakos"
	phraseGerman               string = "nerv"
	phraseItalian              string = "paperonz"
	phraseJapanese             string = "ぎゃくに"
	phraseLatinAmericanSpanish string = "picodoro"
	phraseKorean               string = "유～후"
	phraseRussian              string = "не кряк ли"
	phraseSpanish              string = "picodoro"
	phraseSimplifiedChinese    string = "相反"
	phraseTraditionalChinese   string = "相反"
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
