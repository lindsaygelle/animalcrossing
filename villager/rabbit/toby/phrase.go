package toby

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ribbit"
	phraseCanadianFrench       string = "kérocarot"
	phraseDutch                string = "kikk'r"
	phraseFrench               string = "kérocarot"
	phraseGerman               string = "kerokero"
	phraseItalian              string = "kerokero"
	phraseJapanese             string = "けろけろ"
	phraseLatinAmericanSpanish string = "keroppi"
	phraseKorean               string = "케로케로"
	phraseRussian              string = "зайцеквак"
	phraseSpanish              string = "keroppi"
	phraseSimplifiedChinese    string = "蛙蛙"
	phraseTraditionalChinese   string = "蛙蛙"
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
