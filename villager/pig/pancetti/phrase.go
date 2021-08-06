package pancetti

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sooey"
	phraseCanadianFrench       string = "sabots"
	phraseDutch                string = "kom kom"
	phraseFrench               string = "sabots"
	phraseGerman               string = "grubbel"
	phraseItalian              string = "grufidù"
	phraseJapanese             string = "やだも～"
	phraseLatinAmericanSpanish string = "cuinoink"
	phraseKorean               string = "어우야"
	phraseRussian              string = "хрюшка"
	phraseSpanish              string = "morrito"
	phraseSimplifiedChinese    string = "讨厌啦"
	phraseTraditionalChinese   string = "討厭啦"
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
