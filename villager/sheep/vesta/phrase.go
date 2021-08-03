package vesta

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish    string = ""
	phraseFrench             string = ""
	phraseGerman             string = ""
	phraseItalian            string = ""
	phraseJapanese           string = ""
	phraseKorean             string = ""
	phraseRussian            string = ""
	phraseSpanish            string = ""
	phraseSimplifiedChinese  string = ""
	phraseTraditionalChinese string = ""
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:    phraseAmericanEnglish,
		language.French:             phraseFrench,
		language.German:             phraseGerman,
		language.Italian:            phraseItalian,
		language.Japanese:           phraseJapanese,
		language.Korean:             phraseKorean,
		language.Russian:            phraseRussian,
		language.Spanish:            phraseSpanish,
		language.SimplifiedChinese:  phraseSimplifiedChinese,
		language.TraditionalChinese: phraseTraditionalChinese}
)
