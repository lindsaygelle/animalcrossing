package epona

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "neigh"
	phraseCanadianFrench       string = "hippépique"
	phraseDutch                string = ""
	phraseFrench               string = "hippépique"
	phraseGerman               string = "lonlon"
	phraseItalian              string = "liiinnn"
	phraseJapanese             string = "ヒンヒン"
	phraseLatinAmericanSpanish string = "lon lon"
	phraseKorean               string = "히힝"
	phraseRussian              string = ""
	phraseSpanish              string = "lon lon"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
