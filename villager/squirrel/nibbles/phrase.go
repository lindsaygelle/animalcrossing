package nibbles

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "niblet"
	phraseCanadianFrench       string = "grignotte"
	phraseDutch                string = "knabbel"
	phraseFrench               string = "nunuche"
	phraseGerman               string = "knusper"
	phraseItalian              string = "truciolo"
	phraseJapanese             string = "ガジガジ"
	phraseLatinAmericanSpanish string = "trocitrí"
	phraseKorean               string = "송송"
	phraseRussian              string = "орешек"
	phraseSpanish              string = "bellotita"
	phraseSimplifiedChinese    string = "咯吱咯吱"
	phraseTraditionalChinese   string = "咯吱咯吱"
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
