package cherry

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "what what"
	phraseCanadianFrench       string = "niniche"
	phraseDutch                string = "wat jij"
	phraseFrench               string = "niniche"
	phraseGerman               string = "wedel"
	phraseItalian              string = "arf arf"
	phraseJapanese             string = "だってネ"
	phraseLatinAmericanSpanish string = "guauchi"
	phraseKorean               string = "그냥저냥"
	phraseRussian              string = "что-что"
	phraseSpanish              string = "guauchi"
	phraseSimplifiedChinese    string = "听说娜"
	phraseTraditionalChinese   string = "聽說娜"
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
