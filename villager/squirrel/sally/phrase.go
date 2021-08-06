package sally

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nutmeg"
	phraseCanadianFrench       string = "tac"
	phraseDutch                string = "muskaatje"
	phraseFrench               string = "tac"
	phraseGerman               string = "muskat"
	phraseItalian              string = "nocciolina"
	phraseJapanese             string = "ったら"
	phraseLatinAmericanSpanish string = "raminí"
	phraseKorean               string = "까꿍"
	phraseRussian              string = "мускатик"
	phraseSpanish              string = "ramita"
	phraseSimplifiedChinese    string = "要是"
	phraseTraditionalChinese   string = "要是"
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
