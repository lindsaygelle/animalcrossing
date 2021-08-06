package naomi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "moolah"
	phraseCanadianFrench       string = "merengue"
	phraseDutch                string = "loei"
	phraseFrench               string = "pistouille"
	phraseGerman               string = "joggi"
	phraseItalian              string = "moumou"
	phraseJapanese             string = "セボーン"
	phraseLatinAmericanSpanish string = "mumumu"
	phraseKorean               string = "메르시"
	phraseRussian              string = "муля"
	phraseSpanish              string = "merengue"
	phraseSimplifiedChinese    string = "这样很好"
	phraseTraditionalChinese   string = "這樣很好"
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
