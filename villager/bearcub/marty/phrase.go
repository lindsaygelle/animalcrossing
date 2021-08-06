package marty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pompom"
	phraseCanadianFrench       string = "pom pom"
	phraseDutch                string = "pommetje"
	phraseFrench               string = "pom pom"
	phraseGerman               string = "pompudding"
	phraseItalian              string = "pom pom"
	phraseJapanese             string = "ポムッ"
	phraseLatinAmericanSpanish string = "pudin"
	phraseKorean               string = "폼폼"
	phraseRussian              string = "помпончик"
	phraseSpanish              string = "pudin"
	phraseSimplifiedChinese    string = "布丁"
	phraseTraditionalChinese   string = "布丁"
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
