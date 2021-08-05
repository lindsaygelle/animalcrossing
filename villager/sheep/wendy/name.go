package wendy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Wendy"
	nameCanadianFrench       string = "Karen"
	nameDutch                string = "Wendy"
	nameFrench               string = "Karen"
	nameGerman               string = "Wolli"
	nameItalian              string = "Agnola"
	nameJapanese             string = "みぞれ"
	nameLatinAmericanSpanish string = "Franela"
	nameKorean               string = "눈송이"
	nameRussian              string = "Венди"
	nameSpanish              string = "Franela"
	nameSimplifiedChinese    string = "雪花"
	nameTraditionalChinese   string = "雪花"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
