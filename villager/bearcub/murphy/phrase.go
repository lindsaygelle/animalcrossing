package murphy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "malarkey"
	phraseCanadianFrench       string = "fait que"
	phraseDutch                string = "spruit"
	phraseFrench               string = "canaille"
	phraseGerman               string = "liebelein"
	phraseItalian              string = "grizzel"
	phraseJapanese             string = "ですばい"
	phraseLatinAmericanSpanish string = "heee"
	phraseKorean               string = "샐리"
	phraseRussian              string = "товарищ"
	phraseSpanish              string = "pelusita"
	phraseSimplifiedChinese    string = "罢了"
	phraseTraditionalChinese   string = "罷了"
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
