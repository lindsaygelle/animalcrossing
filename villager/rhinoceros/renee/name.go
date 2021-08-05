package renee

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Renee"
	nameCanadianFrench       string = "Rina"
	nameDutch                string = "Renée"
	nameFrench               string = "Rina"
	nameGerman               string = "Ilona"
	nameItalian              string = "Renata"
	nameJapanese             string = "おさい"
	nameLatinAmericanSpanish string = "Rina"
	nameKorean               string = "뿔님이"
	nameRussian              string = "Рене"
	nameSpanish              string = "Rina"
	nameSimplifiedChinese    string = "柴姐"
	nameTraditionalChinese   string = "柴姐"
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
