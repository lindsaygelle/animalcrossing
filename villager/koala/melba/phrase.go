package melba

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "toasty"
	phraseCanadianFrench       string = "pêchou"
	phraseDutch                string = "toastje"
	phraseFrench               string = "pêchou"
	phraseGerman               string = "zubba"
	phraseItalian              string = "toasty"
	phraseJapanese             string = "とっても"
	phraseLatinAmericanSpanish string = "pichú"
	phraseKorean               string = "정말로요"
	phraseRussian              string = "хрустяшка"
	phraseSpanish              string = "tesoro"
	phraseSimplifiedChinese    string = "不得了"
	phraseTraditionalChinese   string = "不得了"
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
