package dom

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "indeedaroo"
	phraseCanadianFrench       string = "broutille"
	phraseDutch                string = "precies"
	phraseFrench               string = "broutille"
	phraseGerman               string = "jau"
	phraseItalian              string = "veronò"
	phraseJapanese             string = "ふんふん"
	phraseLatinAmericanSpanish string = "digoyó"
	phraseKorean               string = "뿜뿜"
	phraseRussian              string = "точно-точно"
	phraseSpanish              string = "diiigo yo"
	phraseSimplifiedChinese    string = "哇耶"
	phraseTraditionalChinese   string = "哇耶"
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
