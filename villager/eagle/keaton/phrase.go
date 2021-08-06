package keaton

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "wingo"
	phraseCanadianFrench       string = "royaaaal"
	phraseDutch                string = "vleugels"
	phraseFrench               string = "royaaaal"
	phraseGerman               string = "kiiiiiih"
	phraseItalian              string = "baila"
	phraseJapanese             string = "でーッス"
	phraseLatinAmericanSpanish string = "flopflop"
	phraseKorean               string = "수리수리"
	phraseRussian              string = "крыло"
	phraseSpanish              string = "flopflop"
	phraseSimplifiedChinese    string = "是"
	phraseTraditionalChinese   string = "是"
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
