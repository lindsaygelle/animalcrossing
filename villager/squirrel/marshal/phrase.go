package marshal

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sulky"
	phraseCanadianFrench       string = "carpe diem"
	phraseDutch                string = "knorrepot"
	phraseFrench               string = "nanache"
	phraseGerman               string = "huschhusch"
	phraseItalian              string = "ardillà"
	phraseJapanese             string = "あながち"
	phraseLatinAmericanSpanish string = "tecuén"
	phraseKorean               string = "어차피"
	phraseRussian              string = "тихоня"
	phraseSpanish              string = "tecuén"
	phraseSimplifiedChinese    string = "不管怎样"
	phraseTraditionalChinese   string = "不管怎樣"
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
