package becky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chicklet"
	phraseCanadianFrench       string = "côôôôt"
	phraseDutch                string = "kuiken"
	phraseFrench               string = "côôôôt"
	phraseGerman               string = "bokbok"
	phraseItalian              string = "coccoricò"
	phraseJapanese             string = "ハレルヤ"
	phraseLatinAmericanSpanish string = "cucurrús"
	phraseKorean               string = "도레미"
	phraseRussian              string = "модненько"
	phraseSpanish              string = "cucurrús"
	phraseSimplifiedChinese    string = "哈雷路亚"
	phraseTraditionalChinese   string = "哈雷路亞"
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
