package penelope

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "oh bow"
	phraseCanadianFrench       string = "grugru"
	phraseDutch                string = "strikje"
	phraseFrench               string = "grugru"
	phraseGerman               string = "bling"
	phraseItalian              string = "cheeese"
	phraseJapanese             string = "でちゅの"
	phraseLatinAmericanSpanish string = "brillilí"
	phraseKorean               string = "그러찍"
	phraseRussian              string = "бантик"
	phraseSpanish              string = "migas"
	phraseSimplifiedChinese    string = "啾啾"
	phraseTraditionalChinese   string = "啾啾"
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
