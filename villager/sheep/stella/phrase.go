package stella

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "baa-dabing"
	phraseCanadianFrench       string = "mon agneau"
	phraseDutch                string = "mèèèh"
	phraseFrench               string = "mon agneau"
	phraseGerman               string = "badabing"
	phraseItalian              string = "beeene"
	phraseJapanese             string = "ウール"
	phraseLatinAmericanSpanish string = "baa-bii"
	phraseKorean               string = "울"
	phraseRussian              string = "бэ-дабинь"
	phraseSpanish              string = "no veeeas"
	phraseSimplifiedChinese    string = "羊毛"
	phraseTraditionalChinese   string = "羊毛"
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
