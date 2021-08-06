package curly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nyoink"
	phraseCanadianFrench       string = "coink"
	phraseDutch                string = "hupsaknor"
	phraseFrench               string = "coink"
	phraseGerman               string = "oinkoink"
	phraseItalian              string = "niooink"
	phraseJapanese             string = "どもども"
	phraseLatinAmericanSpanish string = "ñoink"
	phraseKorean               string = "꿀꿀"
	phraseRussian              string = "хряк-бряк"
	phraseSpanish              string = "ñoink"
	phraseSimplifiedChinese    string = "感谢"
	phraseTraditionalChinese   string = "感謝"
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
