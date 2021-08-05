package prince

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Prince"
	nameCanadianFrench       string = "Prince"
	nameDutch                string = "Prince"
	nameFrench               string = "Prince"
	nameGerman               string = "Prinz"
	nameItalian              string = "Principe"
	nameJapanese             string = "カール"
	nameLatinAmericanSpanish string = "Felipe"
	nameKorean               string = "카일"
	nameRussian              string = "Принс"
	nameSpanish              string = "Felipe"
	nameSimplifiedChinese    string = "青呱"
	nameTraditionalChinese   string = "青呱"
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
