package cookie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "arfer"
	phraseCanadianFrench       string = "woufi"
	phraseDutch                string = "kwispel"
	phraseFrench               string = "woufi"
	phraseGerman               string = "nuffnuff"
	phraseItalian              string = "bubabù"
	phraseJapanese             string = "プリリン"
	phraseLatinAmericanSpanish string = "can"
	phraseKorean               string = "초롱초롱"
	phraseRussian              string = "гав-гав"
	phraseSpanish              string = "fresita"
	phraseSimplifiedChinese    string = "琳琳"
	phraseTraditionalChinese   string = "琳琳"
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
