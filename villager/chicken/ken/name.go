package ken

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ken"
	nameCanadianFrench       string = "Ken"
	nameDutch                string = "Ken"
	nameFrench               string = "Ken"
	nameGerman               string = "Hannes"
	nameItalian              string = "Aligalli"
	nameJapanese             string = "クロベエ"
	nameLatinAmericanSpanish string = "Pichón"
	nameKorean               string = "오골"
	nameRussian              string = "Кен"
	nameSpanish              string = "Pichón"
	nameSimplifiedChinese    string = "乌骨鸡"
	nameTraditionalChinese   string = "烏骨雞"
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
