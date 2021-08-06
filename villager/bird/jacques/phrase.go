package jacques

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zut alors"
	phraseCanadianFrench       string = "why not"
	phraseDutch                string = "zut alors"
	phraseFrench               string = "ploc ploc"
	phraseGerman               string = "piepiep"
	phraseItalian              string = "pio pio"
	phraseJapanese             string = "チェケラ"
	phraseLatinAmericanSpanish string = "chapó"
	phraseKorean               string = "체키라웃"
	phraseRussian              string = "фу ты"
	phraseSpanish              string = "chapó"
	phraseSimplifiedChinese    string = "唷唷唷"
	phraseTraditionalChinese   string = "唷唷唷"
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
