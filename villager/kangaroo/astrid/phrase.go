package astrid

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "my pet"
	phraseCanadianFrench       string = "mon étoile"
	phraseDutch                string = "troetel"
	phraseFrench               string = "mon étoile"
	phraseGerman               string = "schmusi"
	phraseItalian              string = "fagotto"
	phraseJapanese             string = "だパンク"
	phraseLatinAmericanSpanish string = "brillabrí"
	phraseKorean               string = "뭘봐"
	phraseRussian              string = "лапочка"
	phraseSpanish              string = "mascotita"
	phraseSimplifiedChinese    string = "庞克"
	phraseTraditionalChinese   string = "龐克"
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
