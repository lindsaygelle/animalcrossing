package ketchup

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ketchup"
	nameCanadianFrench       string = "Ketchup"
	nameDutch                string = "Ketchup"
	nameFrench               string = "Ketchup"
	nameGerman               string = "Pullunda"
	nameItalian              string = "Ketchup"
	nameJapanese             string = "ケチャップ"
	nameLatinAmericanSpanish string = "Kétchup"
	nameKorean               string = "케첩"
	nameRussian              string = "Кетчуп"
	nameSpanish              string = "Kétchup"
	nameSimplifiedChinese    string = "番茄酱"
	nameTraditionalChinese   string = "番茄醬"
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
