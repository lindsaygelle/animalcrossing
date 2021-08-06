package felicity

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mimimi"
	phraseCanadianFrench       string = "chatounet"
	phraseDutch                string = "mimimi"
	phraseFrench               string = "chatounet"
	phraseGerman               string = "mauzi"
	phraseItalian              string = "meeeow"
	phraseJapanese             string = "ね"
	phraseLatinAmericanSpanish string = "miauito"
	phraseKorean               string = "우앙"
	phraseRussian              string = "ми-ми-ми"
	phraseSpanish              string = "miauito"
	phraseSimplifiedChinese    string = "哪"
	phraseTraditionalChinese   string = "哪"
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
