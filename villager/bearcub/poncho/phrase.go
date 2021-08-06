package poncho

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "li'l bear"
	phraseCanadianFrench       string = "nounours"
	phraseDutch                string = "beertje"
	phraseFrench               string = "nounours"
	phraseGerman               string = "bärli"
	phraseItalian              string = "babà"
	phraseJapanese             string = "モン"
	phraseLatinAmericanSpanish string = "pruah"
	phraseKorean               string = "몽"
	phraseRussian              string = "медвежонок"
	phraseSpanish              string = "bolita"
	phraseSimplifiedChinese    string = "萌"
	phraseTraditionalChinese   string = "萌"
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
