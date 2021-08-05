package nana

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Nana"
	nameCanadianFrench       string = "Mireille"
	nameDutch                string = "Nana"
	nameFrench               string = "Mireille"
	nameGerman               string = "Dorothea"
	nameItalian              string = "Nanà"
	nameJapanese             string = "チッチ"
	nameLatinAmericanSpanish string = "Nana"
	nameKorean               string = "키키"
	nameRussian              string = "Нана"
	nameSpanish              string = "Nana"
	nameSimplifiedChinese    string = "七七"
	nameTraditionalChinese   string = "七七"
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
