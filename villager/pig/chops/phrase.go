package chops

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zoink"
	phraseCanadianFrench       string = "lardon"
	phraseDutch                string = "moddervet"
	phraseFrench               string = "lardon"
	phraseGerman               string = "borsti"
	phraseItalian              string = "oinc"
	phraseJapanese             string = "だトン"
	phraseLatinAmericanSpanish string = "zoink"
	phraseKorean               string = "떼끼에로"
	phraseRussian              string = "хряк-с"
	phraseSpanish              string = "zoink"
	phraseSimplifiedChinese    string = "豚"
	phraseTraditionalChinese   string = "豚"
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
