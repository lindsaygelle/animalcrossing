package alfonso

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "it'sa me"
	phraseCanadianFrench       string = "moi moi"
	phraseDutch                string = "it⁠'⁠sa me"
	phraseFrench               string = "moi moi"
	phraseGerman               string = "ahhhrg"
	phraseItalian              string = "proprio"
	phraseJapanese             string = "だワニ"
	phraseLatinAmericanSpanish string = "ñam-ñam"
	phraseKorean               string = "나아거"
	phraseRussian              string = "это я"
	phraseSpanish              string = "ñam-ñam"
	phraseSimplifiedChinese    string = "鳄泥"
	phraseTraditionalChinese   string = "鱷泥"
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
