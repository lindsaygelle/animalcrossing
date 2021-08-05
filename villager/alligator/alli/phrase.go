package alli

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish    string = "graaagh"
	phraseFrench             string = "graaagh"
	phraseGerman             string = "graaach"
	phraseItalian            string = "graaag"
	phraseJapanese           string = "どすえ"
	phraseKorean             string = "얘야"
	phraseRussian            string = "гра-а-а"
	phraseSpanish            string = "lagartija"
	phraseSimplifiedChinese  string = "鳄鱼皮"
	phraseTraditionalChinese string = "鱷魚皮"
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
