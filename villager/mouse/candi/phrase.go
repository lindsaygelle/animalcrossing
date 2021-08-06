package candi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sweetie"
	phraseCanadianFrench       string = "claquos"
	phraseDutch                string = "zoetje"
	phraseFrench               string = "claquos"
	phraseGerman               string = "schatzi"
	phraseItalian              string = "squitto"
	phraseJapanese             string = "ですワ"
	phraseLatinAmericanSpanish string = "escuic"
	phraseKorean               string = "달짝지근"
	phraseRussian              string = "сладость"
	phraseSpanish              string = "yogurín"
	phraseSimplifiedChinese    string = "滑滑"
	phraseTraditionalChinese   string = "滑滑"
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
