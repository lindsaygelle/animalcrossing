package gonzo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mate"
	phraseCanadianFrench       string = "calyptus"
	phraseDutch                string = "partner"
	phraseFrench               string = "galopin"
	phraseGerman               string = "partner"
	phraseItalian              string = "dinamite"
	phraseJapanese             string = "だがや"
	phraseLatinAmericanSpanish string = "jorf-jorf"
	phraseKorean               string = "버텨"
	phraseRussian              string = "напарник"
	phraseSpanish              string = "jorf-jorf"
	phraseSimplifiedChinese    string = "钱钱"
	phraseTraditionalChinese   string = "錢錢"
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
