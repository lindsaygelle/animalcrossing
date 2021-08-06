package doc

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "old bunny"
	phraseCanadianFrench       string = "tout bleu"
	phraseDutch                string = "ouwe reus"
	phraseFrench               string = "morbleu"
	phraseGerman               string = "hasi"
	phraseItalian              string = "soufflé"
	phraseJapanese             string = "ですね"
	phraseLatinAmericanSpanish string = "bingboing"
	phraseKorean               string = "그렇잖아"
	phraseRussian              string = "зайчище"
	phraseSpanish              string = "y tal"
	phraseSimplifiedChinese    string = "是啦"
	phraseTraditionalChinese   string = "是啦"
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
