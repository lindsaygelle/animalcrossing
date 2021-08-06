package boyd

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "uh-oh"
	phraseCanadianFrench       string = "pelage"
	phraseDutch                string = "oh-oh"
	phraseFrench               string = "pelage"
	phraseGerman               string = "äffäff"
	phraseItalian              string = "bam bam"
	phraseJapanese             string = "おうおう"
	phraseLatinAmericanSpanish string = "unga"
	phraseKorean               string = "오우오우"
	phraseRussian              string = "ух-ох"
	phraseSpanish              string = "unga"
	phraseSimplifiedChinese    string = "喔喔"
	phraseTraditionalChinese   string = "喔喔"
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
