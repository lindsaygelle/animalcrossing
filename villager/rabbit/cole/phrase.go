package cole

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "coooooool"
	phraseCanadianFrench       string = "chaaauuud"
	phraseDutch                string = "gaaast"
	phraseFrench               string = "chaaauuud"
	phraseGerman               string = "mümml"
	phraseItalian              string = "niglio"
	phraseJapanese             string = "そうさね"
	phraseLatinAmericanSpanish string = "muyayo"
	phraseKorean               string = "나도나도"
	phraseRussian              string = "зайчелло"
	phraseSpanish              string = "muyayo"
	phraseSimplifiedChinese    string = "就是如此"
	phraseTraditionalChinese   string = "就是如此"
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
