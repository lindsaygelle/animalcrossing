package tammi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chimpy"
	phraseCanadianFrench       string = "booof"
	phraseDutch                string = "chimpie"
	phraseFrench               string = "booof"
	phraseGerman               string = "agaga"
	phraseItalian              string = "scimmietta"
	phraseJapanese             string = "ワオ"
	phraseLatinAmericanSpanish string = "uki-uki"
	phraseKorean               string = "와우"
	phraseRussian              string = "шимпатяга"
	phraseSpanish              string = "uki-uki"
	phraseSimplifiedChinese    string = "哇呜"
	phraseTraditionalChinese   string = "哇嗚"
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
