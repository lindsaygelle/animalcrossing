package carrie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "little one"
	phraseCanadianFrench       string = "mon pit"
	phraseDutch                string = "kleine"
	phraseFrench               string = "poupon"
	phraseGerman               string = "beutel"
	phraseItalian              string = "gnegné"
	phraseJapanese             string = "だフン"
	phraseLatinAmericanSpanish string = "arrorró"
	phraseKorean               string = "오롤로"
	phraseRussian              string = "дитятко"
	phraseSpanish              string = "bebé"
	phraseSimplifiedChinese    string = "亲亲"
	phraseTraditionalChinese   string = "親親"
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
