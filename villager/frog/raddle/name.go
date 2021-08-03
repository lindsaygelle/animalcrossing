package raddle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish    string = "Raddle"
	nameFrench             string = ""
	nameGerman             string = ""
	nameItalian            string = ""
	nameJapanese           string = ""
	nameKorean             string = ""
	nameRussian            string = ""
	nameSpanish            string = ""
	nameSimplifiedChinese  string = ""
	nameTraditionalChinese string = ""
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:    nameAmericanEnglish,
		language.French:             nameFrench,
		language.German:             nameGerman,
		language.Italian:            nameItalian,
		language.Japanese:           nameJapanese,
		language.Korean:             nameKorean,
		language.Russian:            nameRussian,
		language.Spanish:            nameSpanish,
		language.SimplifiedChinese:  nameSimplifiedChinese,
		language.TraditionalChinese: nameTraditionalChinese}
)
