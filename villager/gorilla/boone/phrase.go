package boone

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "baboom"
	phraseCanadianFrench       string = "ma noix"
	phraseDutch                string = "bavi-AAN"
	phraseFrench               string = "ma noix"
	phraseGerman               string = "ugh"
	phraseItalian              string = "baboom"
	phraseJapanese             string = "ウルトラ"
	phraseLatinAmericanSpanish string = "angaua"
	phraseKorean               string = "야임마"
	phraseRussian              string = "бабум"
	phraseSpanish              string = "angaua"
	phraseSimplifiedChinese    string = "这家伙"
	phraseTraditionalChinese   string = "超級"
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
