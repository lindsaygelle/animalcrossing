package maple

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "honey"
	phraseCanadianFrench       string = "chouchou"
	phraseDutch                string = "zoeterd"
	phraseFrench               string = "chouchou"
	phraseGerman               string = "darling"
	phraseItalian              string = "miele"
	phraseJapanese             string = "だベア"
	phraseLatinAmericanSpanish string = "mielmiel"
	phraseKorean               string = "저기요"
	phraseRussian              string = "лапушка"
	phraseSpanish              string = "caramelito"
	phraseSimplifiedChinese    string = "熊仔"
	phraseTraditionalChinese   string = "熊仔"
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
