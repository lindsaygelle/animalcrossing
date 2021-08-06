package velma

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "blih"
	phraseCanadianFrench       string = "tu vois"
	phraseDutch                string = "meh"
	phraseFrench               string = "tu vois"
	phraseGerman               string = "mecker"
	phraseItalian              string = "blip"
	phraseJapanese             string = "ザーマス"
	phraseLatinAmericanSpanish string = "verááás-tú"
	phraseKorean               string = "똑바로해"
	phraseRussian              string = "ме-е-е"
	phraseSpanish              string = "veráaas-tú"
	phraseSimplifiedChinese    string = "起立"
	phraseTraditionalChinese   string = "起立"
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
