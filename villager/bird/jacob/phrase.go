package jacob

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ya feel"
	phraseCanadianFrench       string = "asticô"
	phraseDutch                string = "snappie"
	phraseFrench               string = "asticô"
	phraseGerman               string = "tirilü"
	phraseItalian              string = "ciiipito"
	phraseJapanese             string = "っつーの"
	phraseLatinAmericanSpanish string = "chiurp"
	phraseKorean               string = "그거야"
	phraseRussian              string = "чуррп"
	phraseSpanish              string = "pelusilla"
	phraseSimplifiedChinese    string = "话说"
	phraseTraditionalChinese   string = "話說"
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
