package bonbon

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "deelish"
	phraseCanadianFrench       string = "nananère"
	phraseDutch                string = "heer-lijk"
	phraseFrench               string = "nananère"
	phraseGerman               string = "hüppi"
	phraseItalian              string = "hurrah"
	phraseJapanese             string = "ヤバッ"
	phraseLatinAmericanSpanish string = "oh lalá"
	phraseKorean               string = "어뜨케"
	phraseRussian              string = "вкусняшка"
	phraseSpanish              string = "oh lalá"
	phraseSimplifiedChinese    string = "不妙"
	phraseTraditionalChinese   string = "不妙"
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
