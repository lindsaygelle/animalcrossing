package benedict

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "uh-hoo"
	phraseCanadianFrench       string = "oh oh"
	phraseDutch                string = "eitje"
	phraseFrench               string = "oh oh"
	phraseGerman               string = "putputtt"
	phraseItalian              string = "uh-hoo"
	phraseJapanese             string = "ウヒョー"
	phraseLatinAmericanSpanish string = "cocorico"
	phraseKorean               string = "우힛"
	phraseRussian              string = "о-ко-ко"
	phraseSpanish              string = "cocorico"
	phraseSimplifiedChinese    string = "好棒"
	phraseTraditionalChinese   string = "好棒"
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
