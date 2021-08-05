package kabuki

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kabuki"
	nameCanadianFrench       string = "Kabuki"
	nameDutch                string = "Kabuki"
	nameFrench               string = "Kabuki"
	nameGerman               string = "Kabuki"
	nameItalian              string = "Kabuki"
	nameJapanese             string = "かぶきち"
	nameLatinAmericanSpanish string = "Kabuki"
	nameKorean               string = "가북희"
	nameRussian              string = "Кабуки"
	nameSpanish              string = "Kabuki"
	nameSimplifiedChinese    string = "戈伍纪"
	nameTraditionalChinese   string = "戈伍紀"
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
