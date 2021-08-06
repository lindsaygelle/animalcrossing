package chief

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "harrumph"
	phraseCanadianFrench       string = "mmmph"
	phraseDutch                string = "huil"
	phraseFrench               string = "mmmph"
	phraseGerman               string = "harrumff"
	phraseItalian              string = "arrrgh"
	phraseJapanese             string = "やんか"
	phraseLatinAmericanSpanish string = "ejem"
	phraseKorean               string = "그렇잖여"
	phraseRussian              string = "фыр-р-рф"
	phraseSpanish              string = "ejem"
	phraseSimplifiedChinese    string = "不是嘛"
	phraseTraditionalChinese   string = "不是嘛"
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
