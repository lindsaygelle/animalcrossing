package lolly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lolly"
	nameCanadianFrench       string = "Linette"
	nameDutch                string = "Lolly"
	nameFrench               string = "Linette"
	nameGerman               string = "Feline"
	nameItalian              string = "Maty"
	nameJapanese             string = "ラムネ"
	nameLatinAmericanSpanish string = "Feli"
	nameKorean               string = "사이다"
	nameRussian              string = "Лолли"
	nameSpanish              string = "Feli"
	nameSimplifiedChinese    string = "柠檬娜"
	nameTraditionalChinese   string = "檸檬娜"
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
