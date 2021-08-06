package olive

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sweet pea"
	phraseCanadianFrench       string = "trognon"
	phraseDutch                string = "honingkelkje"
	phraseFrench               string = "trognon"
	phraseGerman               string = "honigtopf"
	phraseItalian              string = "fagiolino"
	phraseJapanese             string = "マグ"
	phraseLatinAmericanSpanish string = "panipof"
	phraseKorean               string = "오잉"
	phraseRussian              string = "милашка"
	phraseSpanish              string = "panalito"
	phraseSimplifiedChinese    string = "马克杯"
	phraseTraditionalChinese   string = "馬克杯"
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
