package audie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Audie"
	nameCanadianFrench       string = "Monica"
	nameDutch                string = "Audie"
	nameFrench               string = "Monica"
	nameGerman               string = "Katharina"
	nameItalian              string = "Lupilia"
	nameJapanese             string = "モニカ"
	nameLatinAmericanSpanish string = "Mónica"
	nameKorean               string = "모니카"
	nameRussian              string = "Оди"
	nameSpanish              string = "Mónica"
	nameSimplifiedChinese    string = "莫妮卡"
	nameTraditionalChinese   string = "莫妮卡"
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
