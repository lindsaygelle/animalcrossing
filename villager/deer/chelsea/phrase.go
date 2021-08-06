package chelsea

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pound cake"
	phraseCanadianFrench       string = "bibiche"
	phraseDutch                string = "melodietje"
	phraseFrench               string = "bibiche"
	phraseGerman               string = "denkpink"
	phraseItalian              string = "melodì"
	phraseJapanese             string = "メロメロ"
	phraseLatinAmericanSpanish string = "melomelo"
	phraseKorean               string = "하트하트"
	phraseRussian              string = "бисквит"
	phraseSpanish              string = "melomelo"
	phraseSimplifiedChinese    string = "喜喜"
	phraseTraditionalChinese   string = "喜喜"
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
