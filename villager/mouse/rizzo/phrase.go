package rizzo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "squee"
	phraseCanadianFrench       string = "nouik"
	phraseDutch                string = "piephoi"
	phraseFrench               string = "nouik"
	phraseGerman               string = "fieps"
	phraseItalian              string = "squiii"
	phraseJapanese             string = "がってん"
	phraseLatinAmericanSpanish string = "ñiiiik"
	phraseKorean               string = "힐끔힐끔"
	phraseRussian              string = "пи-и-и"
	phraseSpanish              string = "ñiiiik"
	phraseSimplifiedChinese    string = "蹑蹑"
	phraseTraditionalChinese   string = "躡躡"
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
