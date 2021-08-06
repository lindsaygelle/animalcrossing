package bertha

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bloop"
	phraseCanadianFrench       string = "gloups"
	phraseDutch                string = "blubber"
	phraseFrench               string = "rololol"
	phraseGerman               string = "blubber"
	phraseItalian              string = "blup"
	phraseJapanese             string = "そうです"
	phraseLatinAmericanSpanish string = "glub-glub"
	phraseKorean               string = "맞습니다"
	phraseRussian              string = "бултых"
	phraseSpanish              string = "glub-glub"
	phraseSimplifiedChinese    string = "没错"
	phraseTraditionalChinese   string = "沒錯"
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
