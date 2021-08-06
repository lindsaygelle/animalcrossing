package boris

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "schnort"
	phraseCanadianFrench       string = "minot"
	phraseDutch                string = "oink"
	phraseFrench               string = "minot"
	phraseGerman               string = "quiek"
	phraseItalian              string = "sgrunf"
	phraseJapanese             string = "ブヒッ"
	phraseLatinAmericanSpanish string = "oink"
	phraseKorean               string = "쌀"
	phraseRussian              string = "хрю-хряк"
	phraseSpanish              string = "oink"
	phraseSimplifiedChinese    string = "噗嘻"
	phraseTraditionalChinese   string = "噗嘻"
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
