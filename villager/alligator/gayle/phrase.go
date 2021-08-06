package gayle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snacky"
	phraseCanadianFrench       string = "croc-croc"
	phraseDutch                string = "snackje"
	phraseFrench               string = "croquignou"
	phraseGerman               string = "zähni"
	phraseItalian              string = "désolée"
	phraseJapanese             string = "ワニャン"
	phraseLatinAmericanSpanish string = "smuack"
	phraseKorean               string = "아거얌"
	phraseRussian              string = "перекус"
	phraseSpanish              string = "corazón"
	phraseSimplifiedChinese    string = "鳄莉"
	phraseTraditionalChinese   string = "鱷莉"
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
