package bill

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bill"
	nameCanadianFrench       string = "Choco"
	nameDutch                string = "Bill"
	nameFrench               string = "Choco"
	nameGerman               string = "Bill"
	nameItalian              string = "Gino"
	nameJapanese             string = "ピータン"
	nameLatinAmericanSpanish string = "Paquito"
	nameKorean               string = "코코아"
	nameRussian              string = "Билл"
	nameSpanish              string = "Paquito"
	nameSimplifiedChinese    string = "皮蛋"
	nameTraditionalChinese   string = "皮蛋"
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
