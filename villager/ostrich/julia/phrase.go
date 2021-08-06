package julia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "dahling"
	phraseCanadianFrench       string = "trutruche"
	phraseDutch                string = "schatje"
	phraseFrench               string = "trutruche"
	phraseGerman               string = "püh"
	phraseItalian              string = "blaaah"
	phraseJapanese             string = "やだわね"
	phraseLatinAmericanSpanish string = "churri"
	phraseKorean               string = "어머몰라"
	phraseRussian              string = "дарлинг"
	phraseSpanish              string = "churri"
	phraseSimplifiedChinese    string = "吼唷"
	phraseTraditionalChinese   string = "吼唷"
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
