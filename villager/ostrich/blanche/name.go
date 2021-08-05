package blanche

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Blanche"
	nameCanadianFrench       string = "Sophie"
	nameDutch                string = "Blanche"
	nameFrench               string = "Sophie"
	nameGerman               string = "Christa"
	nameItalian              string = "Emilia"
	nameJapanese             string = "しのぶ"
	nameLatinAmericanSpanish string = "Rocío"
	nameKorean               string = "신옥"
	nameRussian              string = "Бланш"
	nameSpanish              string = "Rocío"
	nameSimplifiedChinese    string = "小偲"
	nameTraditionalChinese   string = "小偲"
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
