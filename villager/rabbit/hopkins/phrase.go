package hopkins

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "thumper"
	phraseCanadianFrench       string = "pfui"
	phraseDutch                string = "stamper"
	phraseFrench               string = "pfui"
	phraseGerman               string = "puschel"
	phraseItalian              string = "siiigh"
	phraseJapanese             string = "ぷぅ"
	phraseLatinAmericanSpanish string = "fiiiiú"
	phraseKorean               string = "뿌우"
	phraseRussian              string = "топ-топ"
	phraseSpanish              string = "vida extra"
	phraseSimplifiedChinese    string = "风"
	phraseTraditionalChinese   string = "風"
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
