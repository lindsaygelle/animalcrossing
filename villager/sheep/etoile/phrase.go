package etoile

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "fuzzy"
	phraseCanadianFrench       string = "mohair"
	phraseDutch                string = "kikilala"
	phraseFrench               string = "mohair"
	phraseGerman               string = "kikilala"
	phraseItalian              string = "ovinin"
	phraseJapanese             string = "メロメロ"
	phraseLatinAmericanSpanish string = "kikilala"
	phraseKorean               string = "폭신폭신"
	phraseRussian              string = "мягенько"
	phraseSpanish              string = "kikilala"
	phraseSimplifiedChinese    string = "毛茸茸"
	phraseTraditionalChinese   string = "毛茸茸"
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
